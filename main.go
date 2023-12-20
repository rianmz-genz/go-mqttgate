package main

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/controller"
	"adriandidimqttgate/exception"
	"adriandidimqttgate/middleware"
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"adriandidimqttgate/seeder"
	"adriandidimqttgate/service"
	"flag"
	"fmt"
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var identityKey = "id"

func helloHandler2(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	session, _ := c.Get(identityKey)
	sessionId := uint(claims["id"].(float64))
	response := web.WebResponse{
		Code:    200,
		Status:  "Success",
		Message: "auth hello successfully",
		Data: map[string]interface{}{
			"sessionId": sessionId,
			"email":     session.(*web.SessionResponse).Email,
			"text":      "Hello World.",
		},
	}
	c.JSON(200, response)
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if rvr := recover(); rvr != nil {
				exception.ErrorHandler(ctx.Writer, ctx.Request, rvr)
			}
		}()
		ctx.Next()
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	r := gin.Default()
	r.Use(ErrorHandler())
	r.Use(CORSMiddleware())

	mqtt := app.NewMqttClient()
	validator := validator.New()
	db := app.NewDBConnection()
	userRepository := repository.NewUserRepository()
	officeRepository := repository.NewOfficeRepository()
	enterActivityRepository := repository.NewEnterActivityRepository()
	sessionRepository := repository.NewSessionRepository()
	authService := service.NewAuthService(userRepository, db, validator)
	qrService := service.NewQrService(enterActivityRepository, officeRepository, sessionRepository, userRepository, db, validator, mqtt)
	authMiddleware := middleware.NewAuthMiddleware(r, db, userRepository, sessionRepository).Middleware()
	authController := controller.NewAuthController(authService)
	qrController := controller.NewQrController(qrService)

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	var tables = []interface{}{
		&domain.Office{},
		&domain.User{},
		&domain.Session{},
		&domain.EnterActivity{},
	}

	fmt.Println("Hello!")
	var m string
	var s string
	flag.StringVar(&m, "m", "none", `migration`)
	flag.StringVar(&s, "s", "none", `seeder`)
	flag.Parse()

	if m == "migrate" {
		db.AutoMigrate(tables...)
	} else if m == "rollback" {
		for i := 0; i < len(tables); i++ {
			db.Migrator().DropTable(tables...)
		}
	} else if m == "refresh" {
		for i := 0; i < len(tables); i++ {
			db.Migrator().DropTable(tables...)
		}
		db.AutoMigrate(tables...)
	} else {
		print("Flag not found")
	}

	if s == "seed" {
		seeder.OfficeSeeder()
		seeder.UserSeeder()
		seeder.SessionSeeder()
	}
	r.GET("/test", func(ctx *gin.Context) {
		test := map[string]interface{}{
			"p": "hai",
		}
		ctx.JSON(200, test)
	})
	r.Use(ErrorHandler())
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", authController.Register)

	auth := r.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler2)
		auth.POST("/logout", authMiddleware.LogoutHandler)
	}

	r.Use(authMiddleware.MiddlewareFunc())
	{
		r.POST("/scan-qr", qrController.ScanQr)
		r.GET("/enter-activities")
	}

	r.Run(":8888")
}
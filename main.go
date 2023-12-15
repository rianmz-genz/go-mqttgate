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
	response := web.WebResponse{
		Code:    200,
		Status:  "Success",
		Message: "auth hello successfully",
		Data: map[string]interface{}{
			"sessionId": claims[identityKey],
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

func main() {
	r := gin.Default()
	r.Use(ErrorHandler())

	validator := validator.New()
	db := app.NewDBConnection()
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()
	authService := service.NewAuthService(userRepository, db, validator)
	authMiddleware := middleware.NewAuthMiddleware(r, db, userRepository, sessionRepository).Middleware()
	authController := controller.NewAuthController(authService)

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

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", authController.Register)

	auth := r.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler2)
		auth.POST("/logout", authMiddleware.LogoutHandler)
	}

	r.Run(":6666")
}

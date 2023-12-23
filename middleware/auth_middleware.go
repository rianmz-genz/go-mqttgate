package middleware

import (
	"adriandidimqttgate/exception"
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var identityKey = "id"

type AuthMiddleware struct {
	Handler           http.Handler
	DB                *gorm.DB
	userRepository    repository.UserRepository
	sessionRepository repository.SessionRepository
}

func NewAuthMiddleware(handler http.Handler, DB *gorm.DB, userRepository repository.UserRepository, sessionRepository repository.SessionRepository) *AuthMiddleware {
	return &AuthMiddleware{
		Handler:           handler,
		DB:                DB,
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}

func (middleware *AuthMiddleware) Middleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "test zone",
		Key:   []byte("secret key"),
		// Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*uint); ok {
				return jwt.MapClaims{
					identityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Println(claims)
			fmt.Println(claims["id"].(float64))
			idUint := uint(claims["id"].(float64))
			session, err := middleware.sessionRepository.GetSessionById(c, middleware.DB, idUint)
			if err != nil {
				panic(exception.NewNotFoundError(err.Error()))
			}

			user := session.User
			return &web.SessionResponse{
				SessionID: session.ID,
				Name:      user.Name,
				Email:     user.Email,
				UserId:    user.ID,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals web.LoginRequest
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userRequest := domain.User{Email: loginVals.Email}
			userResponse, err := middleware.userRepository.GetUserByEmail(c, middleware.DB, userRequest)
			if err != nil {
				return nil, errors.New("user not found")
			}

			if helper.CheckPasswordHash(userRequest.Password, userResponse.Password) {
				return nil, errors.New("wrong password")
			}

			sessionId, err := middleware.sessionRepository.Save(c, middleware.DB, userResponse.ID)
			helper.PanicIfError(err)

			return &sessionId, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// sessionResponse := data.(*web.SessionResponse)

			// middleware.sessionRepository.GetSessionById(c, middleware.DB, sessionResponse.SessionID)

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			response := web.WebResponse{
				Status:  "Fail",
				Code:    http.StatusUnauthorized,
				Message: message,
				Data:    map[string]interface{}{},
			}
			c.JSON(code, response)
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			response := web.WebResponse{
				Status:  "Success",
				Code:    201,
				Message: "login successfully",
				Data: map[string]interface{}{
					"token":  token,
					"expire": t.Format(time.RFC3339),
				},
			}
			c.JSON(http.StatusOK, response)
		},
		LogoutResponse: func(c *gin.Context, code int) {
			// code => http status code => 200
			userResponse, _ := c.Get(identityKey)

			fmt.Println(userResponse)
			sessionId := userResponse.(*web.SessionResponse).SessionID

			_, err := middleware.sessionRepository.DeleteSessionById(c, middleware.DB, sessionId)
			helper.PanicIfError(err)

			response := web.WebResponse{
				Status:  "Success",
				Code:    code,
				Message: "logout successfully",
			}
			c.JSON(http.StatusOK, response)
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

package controller

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService    service.UserService
	SessionService service.SessionService
}

func NewUserController(userService service.UserService, sessionService service.SessionService) UserController {
	return &UserControllerImpl{
		UserService:    userService,
		SessionService: sessionService,
	}
}

func (controller *UserControllerImpl) Profile(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	sessionId := uint(claims["id"].(float64))

	session := controller.SessionService.GetSessionById(c, sessionId)
	userResponse := controller.UserService.GetUserById(c, session.UserId)

	webResponse := web.WebResponse{
		Status:  "Success",
		Code:    200,
		Message: "Get User Profile Successfully",
		Data: map[string]interface{}{
			"user": userResponse,
		},
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

package controller

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/service"

	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	UserService service.AuthService
}

func NewAuthController(userService service.AuthService) AuthController {
	return &AuthControllerImpl{
		UserService: userService,
	}
}

func (controller AuthControllerImpl) Register(c *gin.Context) {
	registerRequest := web.RegisterRequest{}
	helper.ReadFromRequestBody(c.Request, &registerRequest)

	registerResponse := controller.UserService.Register(c, registerRequest)

	webResponse := web.WebResponse{
		Status:  "Success",
		Code:    201,
		Message: "Create new user successfully",
		Data: map[string]interface{}{
			"user": registerResponse,
		},
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

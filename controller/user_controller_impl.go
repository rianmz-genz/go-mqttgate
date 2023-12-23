package controller

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/service"
	"strconv"

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

func (controller UserControllerImpl) Profile(c *gin.Context) {
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

func (controller UserControllerImpl) Update(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	sessionId := uint(claims["id"].(float64))

	result, err := strconv.Atoi(c.Param("userId"))
	helper.PanicIfError(err)
	userId := uint(result)

	updateUserRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(c.Request, &updateUserRequest)

	userUpdateResponse, errUpdate := controller.UserService.Update(c, updateUserRequest, sessionId, userId)
	helper.PanicIfError(errUpdate)

	webResponse := web.WebResponse{
		Status:  "Success",
		Code:    200,
		Message: "Update User Successfully",
		Data: map[string]interface{}{
			"user": userUpdateResponse,
		},
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

func (controller UserControllerImpl) Delete(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	sessionId := uint(claims["id"].(float64))

	result, err := strconv.Atoi(c.Param("userId"))
	helper.PanicIfError(err)
	userId := uint(result)

	controller.UserService.Delete(c, sessionId, userId)

	webResponse := web.WebResponse{
		Status:  "Success",
		Code:    200,
		Message: "Delete User Successfully",
		Data:    map[string]interface{}{},
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

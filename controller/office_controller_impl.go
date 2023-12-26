package controller

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/service"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type OfficeControllerImpl struct {
	OfficeService service.OfficeService
}

func NewOfficeController(officeService service.OfficeService) OfficeController {
	return &OfficeControllerImpl{
		OfficeService: officeService,
	}
}

func (controller OfficeControllerImpl) GetEnterActivitiesByOfficeId(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	sessionId := uint(claims["id"].(float64))

	result, err := strconv.Atoi(c.Param("officeId"))
	helper.PanicIfError(err)
	officeId := uint(result)

	enterActivities, err := controller.OfficeService.GetEntryActivities(c, sessionId, officeId)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Status:  "Success",
		Code:    200,
		Message: "Get Enter Activities By Office Id Successfully",
		Data: map[string]interface{}{
			"enterActivities": enterActivities,
		},
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

func (controller OfficeControllerImpl) CloseGate(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	sessionId := uint(claims["id"].(float64))

	result, err := strconv.Atoi(c.Param("officeId"))
	helper.PanicIfError(err)
	officeId := uint(result)

	closeGateResponse, err := controller.OfficeService.CloseGate(c, sessionId, officeId)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Status:  "Success",
		Code:    200,
		Message: "Close Gate Successfully",
		Data: map[string]interface{}{
			"detail": closeGateResponse,
		},
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}

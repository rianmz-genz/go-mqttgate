package controller

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/service"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type QrControllerImpl struct {
	QrService service.QrService
}

func NewQrController(qrService service.QrService) QrController {
	return &QrControllerImpl{
		QrService: qrService,
	}
}

func (controller QrControllerImpl) ScanQr(c *gin.Context) {
	scanQrRequest := web.ScanQrRequest{}
	helper.ReadFromRequestBody(c.Request, &scanQrRequest)

	claims := jwt.ExtractClaims(c)
	sessionId := uint(claims["id"].(float64))
	
	scanQrResponse := controller.QrService.ScanQr(c, scanQrRequest, sessionId)

	webResponse := web.WebResponse{
		Status: "Success",
		Code: 200,
		Message: "Open Gate Successfully",
		Data: scanQrResponse,
	}

	helper.WriteToResponseBody(c.Writer, webResponse)
}
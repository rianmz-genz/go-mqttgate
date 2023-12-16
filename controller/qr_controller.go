package controller

import "github.com/gin-gonic/gin"

type QrController interface {
	ScanQr(c *gin.Context)
}
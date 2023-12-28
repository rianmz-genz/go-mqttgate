package controller

import "github.com/gin-gonic/gin"

type OfficeController interface {
	GetEnterActivitiesByOfficeId(c *gin.Context)
	CloseGate(c *gin.Context)
	AddOffice(c *gin.Context)
	GetAllOffice(c *gin.Context)
}

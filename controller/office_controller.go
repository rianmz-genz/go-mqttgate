package controller

import "github.com/gin-gonic/gin"

type OfficeController interface {
	GetEnterActivitiesByOfficeId(c *gin.Context)
}

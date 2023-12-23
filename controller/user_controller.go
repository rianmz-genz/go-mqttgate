package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Profile(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

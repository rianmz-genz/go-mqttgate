package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Profile(c *gin.Context)
}

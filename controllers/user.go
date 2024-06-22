package controllers

import (
	"github.com/gin-gonic/gin"
)

func RegisterControllers(r *gin.Engine) {
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	// Add more routes here
}

func GetUsers(c *gin.Context) {
	// Implement your logic here
}

func CreateUser(c *gin.Context) {
	// Implement your logic here
}

package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	// Add more routes here
}

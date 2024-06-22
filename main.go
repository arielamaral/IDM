package main

import (
	"github.com/arielamaral/IDM/controllers"
	"github.com/arielamaral/IDM/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)
	controllers.RegisterControllers(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

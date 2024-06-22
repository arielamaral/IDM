package main

import (
	"github.com/arielamaral/IDM/controllers"
	"github.com/arielamaral/IDM/routes"
	"github.com/arielamaral/IDM/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)
	controllers.RegisterControllers(r)
	services.RegisterServices(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

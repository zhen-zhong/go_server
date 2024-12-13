package main

import (
	"github.com/gin-gonic/gin"
	"learn_go/routes"
	"learn_go/middlewares"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(middlewares.LoggerMiddleware())

	api := r.Group("/api")
	{
		routes.TestoRoutes(api) 
	}
	r.Run(":8080")
}

package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHelloRoutes(rg *gin.RouterGroup) {
	rg.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, API!",
		})
	})
}

package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestoRoutes(rg *gin.RouterGroup) {
	rg.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test, API!",
		})
	})
}

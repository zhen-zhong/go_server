package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// LoggerMiddleware 是一个自定义的日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 继续处理请求
		c.Next()

		// 记录请求日志
		duration := time.Since(start)
		log.Printf("Request %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}

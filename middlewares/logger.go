package middlewares

import (
	"log"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 是一个自定义的日志中间件，日志记录到文件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前日期，用于创建文件夹和文件名
		date := time.Now().Format("2006-01-02") // 格式化为 yyyy-mm-dd
		logDir := "./log/" + date              // 日志文件夹路径
		logFile := logDir + "/log.txt"         // 日志文件路径

		// 创建日志文件夹（如果不存在）
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			err := os.MkdirAll(logDir, os.ModePerm)
			if err != nil {
				log.Printf("Failed to create log directory: %v", err)
				return
			}
		}

		// 打开日志文件（如果文件不存在，创建它；如果存在，追加到文件）
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Failed to open log file: %v", err)
			return
		}
		defer file.Close()

		// 使用 log 包将日志写入文件
		logger := log.New(file, "", log.LstdFlags)

		// 记录请求开始的时间
		start := time.Now()

		// 继续处理请求
		c.Next()

		// 获取请求的状态码
		statusCode := c.Writer.Status()

		// 记录请求日志，包括请求方法、URL、状态码和耗时
		duration := time.Since(start)
		logger.Printf("Request %s %s %d took %v", c.Request.Method, c.Request.URL.Path, statusCode, duration)
	}
}

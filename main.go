package main

import (
	"github.com/gin-gonic/gin"
	"learn_go/routes"
	"learn_go/middlewares"
)

func main() {
	// 创建一个 Gin 引擎实例
	r := gin.Default()

	// 全局中间件
	r.Use(middlewares.LoggerMiddleware())

	// 创建统一的 API 路由组
	api := r.Group("/api")
	{
		// 注册所有接口的路由
		routes.RegisterHelloRoutes(api)  // 注册 hello 接口
	}

	// 启动 HTTP 服务，监听在 8080 端口
	r.Run(":8080")
}

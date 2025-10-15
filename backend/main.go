package main

import (
	"go_gin_mcis/internal/controller"
	"go_gin_mcis/pkg/db"
	"go_gin_mcis/pkg/logger"
	"go_gin_mcis/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志：dev 模式，文件写到 logs/app.log
	logger.Init("logs/app.log")
	// 初始化数据库
	// db.InitOracleDB()
	db.InitDB()

	r := gin.Default()
	// 全局 CORS 中间件
	r.Use(middleware.CORSMiddleware(middleware.Config{
		AllowAll: true, // 开发环境可以改成 true
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://172.31.66.145:3000",
		},
	}))
	// 全局请求日志中间件
	r.Use(middleware.RequestLogger())
	// 注册路由
	controller.InitController(r.Group("/api"))

	logger.Info("server starting on :8081")
	r.Run("0.0.0.0:8081")

}

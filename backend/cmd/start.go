package cmd

import (
	"fmt"
	"go_gin_mcis/bootstrap"
	"go_gin_mcis/config"
	"go_gin_mcis/internal/controller"
	"go_gin_mcis/pkg/logger"
	"go_gin_mcis/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("🔥 cmd.Start panic recovered: %v", r)
		}
	}()

	// 初始化日志（绝对路径）
	logger.Init("logs/app.log")
	r := gin.Default()
	r.Static("/static", "./static")

	// 全局 CORS
	r.Use(middleware.CORSMiddleware(middleware.Config{
		AllowAll: true,
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://172.31.66.145:3000",
		},
	}))
	// 全局请求日志
	r.Use(middleware.RequestLogger())
	controller.InitController(r.Group("/api"))
	bootstrap.StartServices()
	// 优雅关闭 Gin 的 channel
	shutdownChan := make(chan struct{})

	// 启动 Gin
	go func() {
		addr := fmt.Sprintf(":%d", config.GetConf().Server.Port)
		if err := r.Run(addr); err != nil {
			logger.Errorf("❌ Gin 启动失败: %v", err)
		}
		close(shutdownChan)
	}()

	// 阻塞等待服务停止
	<-shutdownChan
}

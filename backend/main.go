package main

import (
	"go_gin_mcis/internal/controller"
	"go_gin_mcis/pkg/db"
	"go_gin_mcis/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志：dev 模式，文件写到 logs/app.log
	logger.Init("logs/app.log")
	// 初始化数据库
	// db.InitDB()

	//初始化oracle
	db.InitOrcalDB()

	r := gin.Default()

	controller.InitController(r.Group("/api"))

	logger.Info("server starting on :8080")
	r.Run(":8080")
}

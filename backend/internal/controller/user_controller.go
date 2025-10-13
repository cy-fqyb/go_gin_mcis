package controller

import (
	"go_gin_mcis/pkg/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterPublicRoutes(func(r *gin.RouterGroup) {
		r.GET("/ping", Ping)
	})
}
func Ping(c *gin.Context) {
	logger.Info("ping endpoint hit")
	c.JSON(200, gin.H{"meassage": "pong"})
}

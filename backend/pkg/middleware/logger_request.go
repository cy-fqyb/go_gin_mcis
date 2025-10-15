// internal/pkg/middleware/logger.go
package middleware

import (
	"go_gin_mcis/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger 打印请求信息中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 请求开始时打印信息
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		logger.Info("Incoming request -> ", method, " ", path, " from ", clientIP)

		// 继续处理请求
		c.Next()

		// 请求结束后打印耗时
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		logger.Info("Completed request -> ", method, " ", path,
			" from ", clientIP,
			" status: ", statusCode,
			" latency: ", latency)
	}
}

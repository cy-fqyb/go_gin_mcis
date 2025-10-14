package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Config 用来配置 CORS 中间件
type Config struct {
	AllowAll       bool
	AllowedOrigins []string
}

// CORSMiddleware 返回一个 Gin 中间件
func CORSMiddleware(cfg Config) gin.HandlerFunc {
	// 把 slice 转 map 方便查找
	originMap := make(map[string]bool)
	for _, o := range cfg.AllowedOrigins {
		originMap[o] = true
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if cfg.AllowAll {
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			if originMap[origin] {
				c.Header("Access-Control-Allow-Origin", origin)
			} else {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Origin not allowed"})
				return
			}
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, X-CSRF-Token")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

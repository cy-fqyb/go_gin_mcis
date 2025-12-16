package middleware

import (
	"go_gin_mcis/pkg/logger"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 你的签名密钥（应放配置文件中）
var jwtSecret = []byte("my_super_secret_key")

// Claims 定义 JWT payload 结构
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

// TokenAuthMiddleware 验证 Token 的中间件
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 里取 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "缺少Authorization头"})
			c.Abort()
			return
		}

		// 一般格式为：Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization格式错误"})
			c.Abort()
			return
		}
		tokenStr := parts[1]

		// 解析 token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			logger.Errorf("无效的token: %v", tokenStr)
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "无效或过期的token"})

			c.Abort()
			return
		}

		// 检查过期时间
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "token已过期"})
			c.Abort()
			return
		}

		// 将用户信息放入上下文，后续接口可直接取
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.Username)

		c.Next()
	}
}

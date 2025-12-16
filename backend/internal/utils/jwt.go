package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my_super_secret_key")

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

// 生成Token
func GenerateToken(userID uint, username string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1天过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cyfqyb",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

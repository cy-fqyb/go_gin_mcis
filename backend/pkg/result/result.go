// internal/pkg/result/result.go
package result

import (
	"encoding/json"
	"go_gin_mcis/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code int    `json:"code"` // 0 表示成功，非 0 表示错误
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, data any) {
	// 日志打印前端响应
	if b, err := json.Marshal(data); err == nil {
		logger.Info("Response data: ", string(b))
	} else {
		logger.Error("Response marshal error: ", err)
	}
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// Fail 返回失败响应
func Fail(c *gin.Context, code int, msg string) {
	logger.Error("Response error: ", msg)
	c.JSON(code, Response{
		Code: code,
		Msg:  msg,
	})
}

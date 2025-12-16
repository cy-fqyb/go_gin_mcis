package controller

import (
	"go_gin_mcis/config"
	"go_gin_mcis/internal/dto"
	"go_gin_mcis/internal/utils"
	"go_gin_mcis/pkg/middleware"
	"go_gin_mcis/pkg/result"
	"reflect"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterPublicRoutes(func(r *gin.RouterGroup) {
		r.POST("/login", middleware.AutoBind(Login, reflect.TypeOf(dto.LoginDTO{})))
		r.GET("/message", func(ctx *gin.Context) {
			result.Success(ctx, gin.H{
				"app_name": config.GetConf().App.Name,
				"version":  config.GetConf().App.Version,
				"tags":     config.GetConf().App.Tags,
			})
		})
	})
}

func Login(c *gin.Context, data any) {
	dto := data.(*dto.LoginDTO)

	if dto.UserName != "cyfqyb" || dto.Password != "123456" {
		result.Fail(c, 401, "用户名或密码错误")
		return
	}

	token, _ := utils.GenerateToken(1, dto.UserName)
	result.Success(c, token)
}

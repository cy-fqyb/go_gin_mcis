package middleware

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"go_gin_mcis/pkg/result"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// HandlerFuncWithDTO 包装器类型
type HandlerFuncWithDTO func(c *gin.Context, dto interface{})

// AutoBind 自动绑定 JSON + 验证 + 类型检查
func AutoBind(handler HandlerFuncWithDTO, dtoType reflect.Type) gin.HandlerFunc {
	return func(c *gin.Context) {
		dto := reflect.New(dtoType).Interface() // 创建 DTO 实例
		if err := c.ShouldBindJSON(dto); err != nil {
			messages := make(map[string]string)

			// 1️⃣ validator 校验错误
			if errs, ok := err.(validator.ValidationErrors); ok {
				val := reflect.TypeOf(dto).Elem()
				for _, e := range errs {
					field, _ := val.FieldByName(e.StructField())
					label := field.Tag.Get("label")
					if label == "" {
						label = e.Field()
					}
					switch e.Tag() {
					case "required":
						messages[e.Field()] = fmt.Sprintf("%s不能为空", label)
					default:
						messages[e.Field()] = fmt.Sprintf("%s不合法", label)
					}
				}
				msg_json, _ := json.Marshal(messages)
				result.Fail(c, 400, string(msg_json))
				c.Abort()
				return
			}

			// 2️⃣ 类型不匹配错误
			errMsg := err.Error()
			if strings.Contains(errMsg, "cannot unmarshal") {
				parts := strings.Split(errMsg, "field ")
				if len(parts) > 1 {
					fieldName := strings.Split(parts[1], " ")[0]
					messages[fieldName] = "参数类型不正确"
					msg_json, _ := json.Marshal(messages)
					result.Fail(c, 400, string(msg_json))
					c.Abort()
					return
				}
			}

			// 3️⃣ 其他 JSON 错误
			result.Fail(c, 400, "JSON 格式错误")
			c.Abort()
			return
		}

		// 绑定成功，调用 handler
		handler(c, dto)
	}
}

package dto

type LoginDTO struct {
	UserName string `json:"user_name" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

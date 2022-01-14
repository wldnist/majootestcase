package dto

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

package dto

type UpdateUserRequest struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
}

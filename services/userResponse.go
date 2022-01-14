package services

import "github.com/wldnist/majootestcase/entities"

type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Token    string `json:"token,omitempty"`
}

func NewUserResponse(user entities.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Name:     user.Name,
	}
}

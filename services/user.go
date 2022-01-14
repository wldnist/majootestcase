package services

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"github.com/wldnist/majootestcase/dto"
	"github.com/wldnist/majootestcase/entities"
	"github.com/wldnist/majootestcase/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(registerRequest dto.RegisterRequest) (*UserResponse, error)
	UpdateUser(updateUserRequest dto.UpdateUserRequest) (*UserResponse, error)
	FindUserByUserName(userName string) (*UserResponse, error)
	FindUserByID(userID string) (*UserResponse, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (c *userService) UpdateUser(updateUserRequest dto.UpdateUserRequest) (*UserResponse, error) {
	user := entities.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&updateUserRequest))

	if err != nil {
		return nil, err
	}

	user, err = c.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	res := NewUserResponse(user)
	return &res, nil

}

func (c *userService) CreateUser(registerRequest dto.RegisterRequest) (*UserResponse, error) {
	user, err := c.userRepository.FindByUserName(registerRequest.UserName)

	if err == nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err = smapping.FillStruct(&user, smapping.MapFields(&registerRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	user, _ = c.userRepository.InsertUser(user)
	res := NewUserResponse(user)
	return &res, nil

}

func (c *userService) FindUserByUserName(userName string) (*UserResponse, error) {
	user, err := c.userRepository.FindByUserName(userName)

	if err != nil {
		return nil, err
	}

	userResponse := NewUserResponse(user)
	return &userResponse, nil
}

func (c *userService) FindUserByID(userID string) (*UserResponse, error) {
	user, err := c.userRepository.FindByUserID(userID)

	if err != nil {
		return nil, err
	}

	userResponse := UserResponse{}
	err = smapping.FillStruct(&userResponse, smapping.MapFields(&user))
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

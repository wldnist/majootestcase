package services

import (
	"errors"
	"log"

	"github.com/wldnist/majootestcase/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(userName string, password string) error
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (c *authService) VerifyCredential(userName string, password string) error {
	user, err := c.userRepo.FindByUserName(userName)
	if err != nil {
		println(err.Error())
		return err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil

}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

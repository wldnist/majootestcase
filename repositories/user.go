package repositories

import (
	"log"

	"github.com/wldnist/majootestcase/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	FindByUserName(userName string) (entities.User, error)
	FindByUserID(userID string) (entities.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{
		connection: connection,
	}
}

func (c *userRepository) InsertUser(user entities.User) (entities.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	c.connection.Save(&user)
	return user, nil
}

func (c *userRepository) UpdateUser(user entities.User) (entities.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entities.User
		c.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	c.connection.Save(&user)
	return user, nil
}

func (c *userRepository) FindByUserName(userName string) (entities.User, error) {
	var user entities.User
	res := c.connection.Where("user_name = ?", userName).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (c *userRepository) FindByUserID(userID string) (entities.User, error) {
	var user entities.User
	res := c.connection.Where("id = ?", userID).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

package repositories

import (
	"github.com/wldnist/majootestcase/entities"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	FindByUserID(userID string) (entities.Merchant, error)
}

type merchantRepository struct {
	connection *gorm.DB
}

func NewMerchantRepository(connection *gorm.DB) MerchantRepository {
	return &merchantRepository{
		connection: connection,
	}
}

func (c *merchantRepository) FindByUserID(userID string) (entities.Merchant, error) {
	var merchant entities.Merchant
	res := c.connection.Where("user_id = ?", userID).Take(&merchant)
	if res.Error != nil {
		return merchant, res.Error
	}
	return merchant, nil
}

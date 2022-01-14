package repositories

import (
	"github.com/wldnist/majootestcase/entities"
	"gorm.io/gorm"
)

type OutletRepository interface {
	All(userID string) ([]entities.Outlet, error)
}

type outletRepository struct {
	connection *gorm.DB
}

func NewOutletRepository(connection *gorm.DB) OutletRepository {
	return &outletRepository{
		connection: connection,
	}
}

func (c *outletRepository) All(userID string) ([]entities.Outlet, error) {
	outlets := []entities.Outlet{}

	var merchant entities.Merchant
	res := c.connection.Where("user_id = ?", userID).Take(&merchant)
	if res.Error != nil {
		return outlets, res.Error
	}

	c.connection.Preload("Merchant").Where("merchant_id = ?", merchant.ID).Find(&outlets)
	return outlets, nil
}

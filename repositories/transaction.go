package repositories

import (
	"github.com/wldnist/majootestcase/entities"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	All(userID string) ([]entities.Transaction, error)
}

type transactionRepository struct {
	connection *gorm.DB
}

func NewTransactionRepository(connection *gorm.DB) TransactionRepository {
	return &transactionRepository{
		connection: connection,
	}
}

func (c *transactionRepository) All(userID string) ([]entities.Transaction, error) {
	transactions := []entities.Transaction{}

	var merchant entities.Merchant
	res := c.connection.Where("user_id = ?", userID).Take(&merchant)
	if res.Error != nil {
		return transactions, res.Error
	}

	var outlet entities.Outlet
	res2 := c.connection.Where("merchant_id = ?", merchant.ID).Take(&outlet)
	if res.Error != nil {
		return transactions, res2.Error
	}

	c.connection.Preload("Outlet").Where("outlet_id = ?", outlet.ID).Find(&transactions)
	return transactions, nil
}

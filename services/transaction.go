package services

import (
	"github.com/wldnist/majootestcase/repositories"
)

type TransactionService interface {
	All(userID string) (*[]TransactionResponse, error)
}

type transactionService struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (c *transactionService) All(userID string) (*[]TransactionResponse, error) {
	transactions, err := c.transactionRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := NewTransactionArrayResponse(transactions)
	return &prods, nil
}

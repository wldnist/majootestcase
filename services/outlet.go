package services

import (
	"github.com/wldnist/majootestcase/repositories"
)

type OutletService interface {
	All(userID string) (*[]OutletResponse, error)
}

type outletService struct {
	outletRepo repositories.OutletRepository
}

func NewOutletService(outletRepo repositories.OutletRepository) OutletService {
	return &outletService{
		outletRepo: outletRepo,
	}
}

func (c *outletService) All(userID string) (*[]OutletResponse, error) {
	outlets, err := c.outletRepo.All(userID)
	if err != nil {
		return nil, err
	}

	prods := NewOutletArrayResponse(outlets)
	return &prods, nil
}

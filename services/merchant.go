package services

import (
	"github.com/mashingan/smapping"
	"github.com/wldnist/majootestcase/repositories"
)

type MerchantService interface {
	FindMerchantByUserID(userID string) (*MerchantResponse, error)
}

type merchantService struct {
	merchantRepo repositories.MerchantRepository
}

func NewMerchantService(merchantRepo repositories.MerchantRepository) MerchantService {
	return &merchantService{
		merchantRepo: merchantRepo,
	}
}

func (c *merchantService) FindMerchantByUserID(userID string) (*MerchantResponse, error) {
	merchants, err := c.merchantRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	merchantResponse := MerchantResponse{}
	err = smapping.FillStruct(&merchantResponse, smapping.MapFields(&merchants))
	if err != nil {
		return nil, err
	}
	return &merchantResponse, nil
}

package services

import (
	"github.com/wldnist/majootestcase/entities"
)

type MerchantResponse struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	MerchantName string `json:"merchant"`
}

func NewMerchantResponse(merchant entities.Merchant) MerchantResponse {
	return MerchantResponse{
		ID:           merchant.ID,
		UserID:       merchant.UserID,
		MerchantName: merchant.MerchantName,
	}
}

func NewMerchantArrayResponse(merchant []entities.Merchant) []MerchantResponse {
	merchantRes := []MerchantResponse{}
	for _, v := range merchant {
		p := MerchantResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			MerchantName: v.MerchantName,
		}
		merchantRes = append(merchantRes, p)
	}
	return merchantRes
}

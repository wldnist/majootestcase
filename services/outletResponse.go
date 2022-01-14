package services

import (
	"github.com/wldnist/majootestcase/entities"
)

type OutletResponse struct {
	ID         int64            `json:"id"`
	OutletName string           `json:"outlet_name"`
	Merchant   MerchantResponse `json:"merchant,omitempty"`
}

func NewOutletResponse(outlet entities.Outlet) OutletResponse {
	return OutletResponse{
		ID:         outlet.ID,
		OutletName: outlet.OutletName,
		Merchant:   NewMerchantResponse(outlet.Merchant),
	}
}

func NewOutletArrayResponse(outlet []entities.Outlet) []OutletResponse {
	outletRes := []OutletResponse{}
	for _, v := range outlet {
		p := OutletResponse{
			ID:         v.ID,
			OutletName: v.OutletName,
			Merchant:   NewMerchantResponse(v.Merchant),
		}
		outletRes = append(outletRes, p)
	}
	return outletRes
}

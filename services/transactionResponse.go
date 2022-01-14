package services

import (
	"github.com/wldnist/majootestcase/entities"
)

type TransactionResponse struct {
	ID        int64          `json:"id"`
	BillTotal float64        `json:"bill_total"`
	Outlet    OutletResponse `json:"outlet,omitempty"`
}

func NewTransactionResponse(transaction entities.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:        transaction.ID,
		BillTotal: transaction.BillTotal,
		Outlet:    NewOutletResponse(transaction.Outlet),
	}
}

func NewTransactionArrayResponse(transaction []entities.Transaction) []TransactionResponse {
	transactionRes := []TransactionResponse{}
	for _, v := range transaction {
		p := TransactionResponse{
			ID:        v.ID,
			BillTotal: v.BillTotal,
			Outlet:    NewOutletResponse(v.Outlet),
		}
		transactionRes = append(transactionRes, p)
	}
	return transactionRes
}

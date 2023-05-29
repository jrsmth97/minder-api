package view

import (
	"minder/src/server/model"
)

type SyncPurchaseResponse struct {
	Id            string `json:"id"`
	InvoiceNumber string `json:"invoice_number"`
	PaymentMethod string `json:"payment_method"`
	PaymentStatus string `json:"payment_status"`
}

type FindPurchaseResponse struct {
	SyncPurchaseResponse
	VANumber string `json:"va_number"`
}

func NewSyncPurchaseResponse(purchases *[]model.Purchase) *[]SyncPurchaseResponse {
	var syncPurchaseResponses []SyncPurchaseResponse

	for _, purchase := range *purchases {
		syncPurchaseResponses = append(syncPurchaseResponses, SyncPurchaseResponse{
			Id:            purchase.ID.String(),
			InvoiceNumber: purchase.InvoiceNumber,
			PaymentMethod: purchase.PaymentMethod,
			PaymentStatus: purchase.Status,
		})
	}

	return &syncPurchaseResponses
}

func NewFindPurchaseResponse(purchase *model.Purchase) *FindPurchaseResponse {
	findPurchase := &FindPurchaseResponse{}
	findPurchase.Id = purchase.ID.String()
	findPurchase.InvoiceNumber = purchase.InvoiceNumber
	findPurchase.PaymentMethod = purchase.PaymentMethod
	findPurchase.PaymentStatus = purchase.Status
	findPurchase.VANumber = purchase.VaNumber

	return findPurchase
}

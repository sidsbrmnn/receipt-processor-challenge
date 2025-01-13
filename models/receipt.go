package models

import (
	"github.com/google/uuid"
)

type Receipt struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Retailer     string    `json:"retailer" binding:"required,regex=^[\\w\\s\\-&]+$"`
	PurchaseDate DateOnly  `json:"purchaseDate" binding:"required"`
	PurchaseTime TimeOnly  `json:"purchaseTime" binding:"required"`
	Items        []struct {
		ShortDescription string `json:"shortDescription" binding:"required,regex=^[\\w\\s\\-]+$"`
		Price            Price  `json:"price" binding:"required,precision=2"`
	} `json:"items" binding:"required,gt=0,dive"`
	Total Price `json:"total" binding:"required,precision=2"`
}

// interface for interacting with the receipt repository
type ReceiptRepository interface {
	Create(Receipt) (string, error)
	GetByID(string) (Receipt, error)
}

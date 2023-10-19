package receipt

import (
	"github.com/google/uuid"
)

type Receipt struct {
	ID           uuid.UUID
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        string
	Items        []Item
}

type Item struct {
	ShortDescription string
	Price            float32
}

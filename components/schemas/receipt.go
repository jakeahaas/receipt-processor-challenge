package receipt

type Receipt struct {
	ID           ID 	`json:"id"`
	Points		 Points	`json:"points"`
	Retailer     string	`json:"retailer"`
	PurchaseDate string	`json:"purchaseDate"`
	PurchaseTime string	`json:"purchaseTime"`
	Total        string	`json:"total"`
	Items        []Item	`json:items"`
}

type Item struct {
	ShortDescription string	`json:"shortDescription"`
	Price            string	`json:"price"`
}

type ID struct {
	ID	string	`json:"id"`
}

type Points struct {
	Points int64	`json:"points"`
}
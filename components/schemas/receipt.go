package receipt

type Receipt struct {
	ID          	ID 			`json:"id"`
	Points		 	Points		`json:"points"`
	Retailer    	string		`json:"retailer" binding:"required"`
	PurchaseDate 	string		`json:"purchaseDate" binding:"required"`
	PurchaseTime 	string		`json:"purchaseTime" binding:"required"`
	Items        	[]Item		`json:"items" binding:"required,dive"`
	Total        	float64		`json:"total,string" binding:"required"`
}

type Item struct {
	ShortDescription string	`json:"shortDescription" binding:"required"`
	Price            float64	`json:"price,string" binding:"required"`
}

type ID struct {
	ID	string	`json:"id"`
}

type Points struct {
	Points int64	`json:"points"`
}
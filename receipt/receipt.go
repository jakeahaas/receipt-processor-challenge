package receipt

type Receipt struct {
	ID           string
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        string
	Items        []Item
}

type Item struct {
	ShortDescription string
	Price            string
}

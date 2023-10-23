package main

import (
	"net/http"
	"github.com/jakeahaas/receipt-processor-challenge/components/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math"
	"strings"
	"time"
)

var receipts = []receipt.Receipt{}

//Main function runs the router to process endpoints
func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", findReceipt)
	router.Run("localhost:8080")
}

//Tries to add the receipt to the receipts array
func processReceipt(c *gin.Context) {
	var newReceipt receipt.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	//Easiest way to check if date and time is in proper format, this isnt actually used at all anywhere
	date, err := time.Parse("2006-01-02", newReceipt.PurchaseDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, "The receipt is invalid (invalid purchaseDate)")
		return
	}
	time, err := time.Parse("15:04", newReceipt.PurchaseTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, "The receipt is invalid (invalid purchaseTime)")
		return
	}
	newReceipt.Date = date
	newReceipt.Time = time
	//End date and time check, again this is not used anywhere else
	////////////////////////////////////////////////////////////////////////////////////////////////////
	
	//Create a new ID for this receipt
	var newID receipt.ID
	newID.ID = uuid.New().String()
	newReceipt.ID.ID = newID.ID

	//Add this receipt to the receipts array
	receipts = append(receipts, newReceipt)
	c.JSON(http.StatusOK, newID)
}

//Tries to find a receipt with the given receipt ID
func findReceipt(c *gin.Context) {
	var id string
	id = c.Params.ByName("id")
	for i := 0; i < len(receipts); i++ {
		if (receipts[i].ID.ID == id) {
			scoreReceipt(receipts[i], c)
			return
		}
	}
	c.JSON(http.StatusNotFound, "No receipt found for that id")
}

//This function scores the receipt that is input.  This is the main logic portion of the code
func scoreReceipt(receipt receipt.Receipt, c *gin.Context) {
	//Check retail name for alphanumeric characters
	for i := 0; i < len(receipt.Retailer); i++ {
		char := receipt.Retailer[i]
		if ('a' <= char && char <= 'z' ||
			'A' <= char && char <= 'Z' ||
			'0' <= char && char <= '9') {
				receipt.Points.Points ++
			}
	}
	//Check if total is a multiple of .25
	var cents float64
	cents = receipt.Total - math.Round(receipt.Total - .5)
	//if amount ends in .00, its a round dollar amount and thus a multiple of .25 (so +50 + 25)
	if (cents == .00) {
		receipt.Points.Points += 75
	}
	if (cents == .25 || cents == .50 || cents == .75) {
		receipt.Points.Points += 25
	}
	//Check how many pairs of items there are
	receipt.Points.Points += int64(5 * (len(receipt.Items)/2))
	//Check if item description length is a multiple of 3
	for i := 0; i < len(receipt.Items); i++ {
		trimmed := strings.TrimSpace(receipt.Items[i].ShortDescription)
		if (len(trimmed) % 3 == 0) {
			receipt.Points.Points += int64(math.Round(( receipt.Items[i].Price * float64(.2) ) + .5))
		}
	}
	// //if purchase day is odd
	tempNum := receipt.PurchaseDate[len(receipt.PurchaseDate) - 1:]
	if (tempNum == "1" || tempNum == "3" || tempNum == "5" || tempNum == "7" || tempNum == "9") {
		receipt.Points.Points += 6
	}
	//check if purchase is after 2:00 PM and before 4:00 PM
	purchaseHour := receipt.PurchaseTime[0:2]
	if purchaseHour == "14" || purchaseHour == "15" {
		receipt.Points.Points += 10
	}
	c.JSON(http.StatusOK, receipt.Points)
}
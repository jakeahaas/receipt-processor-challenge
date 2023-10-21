package main

import (
	"net/http"
	"github.com/jakeahaas/receipt-processor-challenge/components/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"math"
	"strings"
	"strconv"
)

var receipts = []receipt.Receipt{}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReceipt)
	router.GET("/receipts/:id/points", findReceipt)
	router.Run("localhost:8080")
}

func processReceipt(c *gin.Context) {
	var newReceipt receipt.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, "The receipt is invalid")
		return
	}

	var newID receipt.ID
	newID.ID = uuid.New().String()
	newReceipt.ID.ID = newID.ID

	receipts = append(receipts, newReceipt)
	c.JSON(http.StatusOK, newID)
}

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
	var cents string
	cents = receipt.Total[len(receipt.Total) - 3:] ///////////////////////////////////////////////////////
	if (cents == ".25" || cents == ".50" || cents == ".75") {
		receipt.Points.Points += 25
	}
	//if amount ends in .00, its a round dollar amount and thus a multiple of .25 (so +50 + 25)
	if (cents == ".00") {
		receipt.Points.Points += 75
	}
	//Check how many pairs of items there are
	receipt.Points.Points += int64(5 * (len(receipt.Items)/2))
	//Check if item description length is a multiple of 3
	for i := 0; i < len(receipt.Items); i++ {
		trimmed := strings.TrimSpace(receipt.Items[i].ShortDescription)
		if (len(trimmed) % 3 == 0) {
			price, err := strconv.ParseFloat(receipt.Items[i].Price, 64)
			if err != nil {
				//THIS SHOULD BE CHECKED EARLIER
				c.JSON(http.StatusBadRequest, "issue in items array")
				return
			}
			receipt.Points.Points += int64(math.Round(( price * float64(.2) ) + .5))
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
	c.JSON(http.StatusOK, receipt)
}

package main

import (
	"net/http"
	"github.com/RECEIPT-PROCESSOR-CHALLENGE/receipt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = []receipt.Receipt{}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", postReceipt)
	//router.GET("/receipts/{id}/points", scoreReceipt)
	router.GET("/receipts/points", scoreReceipt)
	router.Run("localhost:8080")
}

func postReceipt(c *gin.Context) {
	var newReceipt receipt.Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt.ID = uuid.New()
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

func scoreReceipt(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "receipts")
}

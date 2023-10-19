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
	router.GET("/receipts/:id/points", scoreReceipt)
	//THIS IS TESTING STUFF HERE
	router.GET("/receipts/points", scoreReceipt)
	//END TESTING STUFF HERE
	router.Run("localhost:8080")
}

func postReceipt(c *gin.Context) {
	var newReceipt receipt.Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	newReceipt.ID = uuid.New().String()
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

func scoreReceipt(c *gin.Context) {
	var id string
	id = c.Params.ByName("id")
	for i := 0; i < len(receipts); i++ {
		if receipts[i].ID == id {
			//found this receipt, now do math
			c.IndentedJSON(http.StatusOK, receipts[i])
			break
		}
	}
}

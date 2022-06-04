package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetCoinAvailableQuantity(c *gin.Context) {
	customerId := c.PostForm("customerid")
	coin := c.PostForm("coin")
	log.Println("customerid=", customerId, " coin=", coin)
	c.JSON(200, 100)
}

func ConfirmTransfer(c *gin.Context) {
	orderId := c.PostForm("orderid")
	success := c.PostForm("success")
	log.Println("orderId=", orderId, " success=", success)
	c.String(200, "true")
}

func TransferAssets(c *gin.Context) {
	token := c.PostForm("token")
	orderId := c.PostForm("orderid")
	amount := c.PostForm("amount")
	side := c.PostForm("side")
	coin := c.PostForm("coin")
	log.Println("orderId=", orderId, " token=", token, "amount=", amount, "side=", side, "coin=", coin)
	c.String(200, "true")
}

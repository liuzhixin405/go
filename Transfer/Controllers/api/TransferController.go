package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type TransferController struct{}

func (tra TransferController) GetCoinAvailableQuantity(c *gin.Context) {
	customerId := c.PostForm("customerid")
	coin := c.PostForm("coin")
	log.Println("customerid=", customerId, " coin=", coin)
	c.JSON(200, 100)
}

func (tra TransferController) ConfirmTransfer(c *gin.Context) {
	orderId := c.PostForm("orderid")
	success := c.PostForm("success")
	log.Println("orderId=", orderId, " success=", success)
	c.String(200, "true")
}

func (tra TransferController) TransferAssets(c *gin.Context) {
	token := c.PostForm("token")
	orderId := c.PostForm("orderid")
	amount := c.PostForm("amount")
	side := c.PostForm("side")
	coin := c.PostForm("coin")
	log.Println("orderId=", orderId, " token=", token, "amount=", amount, "side=", side, "coin=", coin)
	c.String(200, "true")
}

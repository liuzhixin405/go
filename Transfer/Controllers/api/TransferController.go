package api

import (
	"github.com/gin-gonic/gin"
)

type TransferController struct{}

func (tra TransferController) GetCoinAvailableQuantity(c *gin.Context) {
	c.JSON(200, 100)
}

func (tra TransferController) ConfirmTransfer(c *gin.Context) {
	c.String(200, "true")
}

func (tra TransferController) TransferAssets(c *gin.Context) {
	c.String(200, "true")
}

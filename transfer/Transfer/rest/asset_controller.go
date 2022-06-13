package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

// CreateResource godoc
// @Summary 获取对方系统可用资产
// @Description 获取对方系统可用资产
// @Tags asset 相关接口
// @Param customerid path string true "用户id"
// @Param coin path string true "币种id"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /getCoinAvailablequantity [post]
func GetCoinAvailableQuantity(c *gin.Context) {
	customerId := c.PostForm("customerid")
	coin := c.PostForm("coin")
	log.Println("customerid=", customerId, " coin=", coin)
	c.JSON(200, 100)
}

// CreateResource godoc
// @Summary 资产划转确认
// @Description 资产划转确认
// @Tags asset 相关接口
// @Param orderid path string true "orderid"
// @Param success path string true "是否成功"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /confirmtransfer [post]
func ConfirmTransfer(c *gin.Context) {
	orderId := c.PostForm("orderid")
	success := c.PostForm("success")
	log.Println("orderId=", orderId, " success=", success)
	c.String(200, "true")
}

// CreateResource godoc
// @Summary  资产划转
// @Description 资产划转
// @Tags asset 相关接口
// @Param token path string true "用户token"
// @Param coin path string true "币种id"
// @Param orderid path string true "orderid"
// @Param amount path string true "转账金额"
// @Param side path string true "转账方向"
// @Accept  json
// @Success 200  {object} string "success"
// @Failure 400  {string} string "error"
// @Failure 404  {string} string "error"
// @Failure 500  {string} string "error"
// @Router /transferassets [post]
func TransferAssets(c *gin.Context) {
	token := c.PostForm("token")
	orderId := c.PostForm("orderid")
	amount := c.PostForm("amount")
	side := c.PostForm("side")
	coin := c.PostForm("coin")
	log.Println("orderId=", orderId, " token=", token, "amount=", amount, "side=", side, "coin=", coin)
	c.String(200, "true")
}

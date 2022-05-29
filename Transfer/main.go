package main

import (
	"net/http"
	"transferasset/Controllers/api"
	"transferasset/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitPsqlDB()

	router := gin.Default()
	v1 := router.Group("/v1/transfer")
	{
		v1.POST("/confirmtransfer", api.TransferController{}.ConfirmTransfer)
		v1.POST("/getavailablequantity", api.TransferController{}.GetCoinAvailableQuantity)
		v1.POST("/transferasset", api.TransferController{}.TransferAssets)
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic("服务启动失败!")
	}
}

//curl -X POST http://localhost:8080/v1/transfer/getavailablequantity  -d coin=1  -d customerId=test001
//curl -X POST http://localhost:8080/v1/transfer/confirmtransfer
// curl -X POST http://localhost:8080/v1/transfer/transferasset

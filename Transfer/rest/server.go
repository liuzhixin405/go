package rest

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	router := gin.Default()
	router.Use(setCROSOptions)

	v1 := router.Group("/v1/transfer")
	{
		v1.POST("/confirmtransfer", ConfirmTransfer)
		v1.POST("/getavailablequantity", GetCoinAvailableQuantity)
		v1.POST("/transferasset", TransferAssets)
	}

	err := router.Run(server.addr)
	if err != nil {
		panic(err)
	}
}

func setCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
}

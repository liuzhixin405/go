package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "transferasset/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router := gin.New()
	router.Use(setCROSOptions)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome To Sample program swagger"})
	})
	v1 := router.Group("/v1/transfer")
	{
		v1.POST("/confirmtransfer", ConfirmTransfer)
		v1.POST("/getavailablequantity", GetCoinAvailableQuantity)
		v1.POST("/transferasset", TransferAssets)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(server.addr)
	//err := router.Run()
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

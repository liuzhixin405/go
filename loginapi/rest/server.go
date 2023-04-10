package rest

import (
	"io/ioutil"
	"net/http"

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
	router := gin.New()
	router.Use(setCROSOptions)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome To Sample program swagger"})
	})
	v1 := router.Group("/v1/login")
	{
		v1.POST("/", Login)
		v1.GET("/greeting", Greeting)
	}
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

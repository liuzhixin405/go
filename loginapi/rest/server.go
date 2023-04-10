package rest

import (
	"io/ioutil"
	"login/service"
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
	isApi := false
	if isApi {
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

	} else {
		htmlrouter := gin.Default()
		htmlrouter.LoadHTMLGlob("templates/*")

		htmlrouter.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		htmlrouter.POST("/login", func(c *gin.Context) {
			password := c.PostForm("password")
			username := c.PostForm("username")

			var lservice service.LoginService
			lservice = service.Service{}
			isLogin := lservice.Login(username, password)
			if isLogin {
				c.HTML(http.StatusOK, "success.html", gin.H{
					"username": username,
				})
			} else {
				c.HTML(http.StatusOK, "error.html", nil)
			}
		})
		err := htmlrouter.Run(server.addr)
		//err := router.Run()
		if err != nil {
			panic(err)
		}
	}
}

func setCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
}

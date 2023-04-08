package rest

import (
	"log"
	"login/service"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry/service"
)

func Login(c *gin.Context) {
	userName := c.PostForm("suername")
	password := c.PostForm("password")
	isLogin := service.Login(userName, password)
	if isLogin != true {
		c.JSON(404, "账号密码错误")
	}

	log.Println("username=", userName, "password=", password)
	c.JSON(200, "string")
}

func Greeting(c *gin.Context) {
	c.JSON(200, "hello")
}

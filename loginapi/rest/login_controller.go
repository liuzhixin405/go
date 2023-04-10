package rest

import (
	"log"
	"login/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginDTO loginDto

	if err := c.BindJSON(&loginDTO); err != nil {
		log.Fatal(err)
	}

	var lservice service.LoginService
	lservice = service.Service{}
	isLogin := lservice.Login(loginDTO.UserName, loginDTO.PassWord)
	if isLogin != true {
		c.JSON(404, "账号密码错误")
		log.Println("username=", loginDTO.UserName, "password=", loginDTO.PassWord)
	} else {
		c.JSON(200, "登陆成功")
	}

}

func Greeting(c *gin.Context) {
	c.JSON(200, "hello")
}

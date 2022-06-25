package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

//Main API
type MainController struct {
	beego.Controller
}

// @Title 测试
// @Description get Main
// @Success 200
// @Failure 400 Invalid email supplied
// @Failure 404 main not found
// @router /get/ [get]
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// func (c *MainController) Get() {
// 	c.Ctx.WriteString("hello")
// }

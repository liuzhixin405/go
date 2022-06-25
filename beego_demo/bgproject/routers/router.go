package routers

import (
	"bgproject/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/main", beego.NSInclude(
			&controllers.MainController{},
		),
		),
	)
	beego.AddNamespace(ns)
	//eego.Router("/", &controllers.MainController{})
}

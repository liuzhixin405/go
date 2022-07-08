package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["bgproject/controllers:MainController"] = append(beego.GlobalControllerRouter["bgproject/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/get/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

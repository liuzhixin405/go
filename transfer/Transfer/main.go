package main

import (
	_ "transferasset/docs"
	"transferasset/rest"
)

// @title 资产划转
// @version 1.0
// @description 不同系统之间账户资产划转
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8111
// @BasePath /v1/transfer
func main() {
	rest.StartServer()

	select {}
}

//curl -X POST http://localhost:8080/v1/transfer/getavailablequantity  -d coin=1  -d customerId=test001
//curl -X POST http://localhost:8080/v1/transfer/confirmtransfer
// curl -X POST http://localhost:8080/v1/transfer/transferasset

// go install github.com/swaggo/swag/cmd/swag@latest 安装swagger需要的doc文档
//执行swag init   (检查docs.go 注意格式)
//http://localhost:8111/swagger/index.html

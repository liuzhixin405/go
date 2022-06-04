package rest

import (
	"fmt"
	"transferasset/conf"
)

func StartServer() {
	gbeConfig := conf.GetConfig()
	httpServer := NewHttpServer(gbeConfig.DataSource.Addr)
	go httpServer.Start()
	fmt.Println("rest servwer start")
}

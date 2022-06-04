package rest

import (
	"log"
	"transferasset/conf"
)

func StartServer() {
	gbeConfig := conf.GetConfig()
	httpServer := NewHttpServer(gbeConfig.DataSource.Addr)
	go httpServer.Start()
	log.Println("rest servwer start")
}

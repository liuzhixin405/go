package rest

import (
	"log"
	"login/conf"
)

func StartServer() {
	gbeCOnfig := conf.GetConfig()
	httpServer := NewHttpServer(gbeCOnfig.RestServer.Addr)
	go httpServer.Start()
	log.Println("rest server start")
}

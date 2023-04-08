package main

import (
	"login/rest"
)

func main() {
	rest.StartServer()
	select {}
}

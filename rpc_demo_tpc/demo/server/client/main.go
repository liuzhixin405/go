package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer conn.Close()

	var reply string
	err2 := conn.Call("hello.SayHello", "我是客户端,服务端在吗", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(reply)
}

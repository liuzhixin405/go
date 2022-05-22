package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//conn, err1 := rpc.Dial("tcp", "192.168.253.130:8080")
	conn, err1 := net.Dial("tcp", "192.168.253.130:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer conn.Close()
	//建立基于json编码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)) //jsonrpc新增
	var reply string
	//err2 := conn.Call("hello.SayHello", "我是客户端,服务端在吗", &reply)
	err2 := client.Call("hello.SayHello", "我是客户端,服务端在吗", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(reply)
}

//需要关闭centos的防火墙 或者开放8080端口  查看 centos的IP 本地测试为:(192.168.253.130)
//执行 nc -l 192.168.253.130 8080  (安装nc  为 yum instal  -y nc)

package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {
}

/*
1、方法只能有两个可以序列化的参数，第二个是指针类型
2、方法返回error，必须是公开方法
3、req、res类型不能是 func channel complex 不能序列化
*/
func (h Hello) SayHello(req string, res *string) error {
	fmt.Println(req)
	*res = "你好, " + req
	return nil
}
func main() {
	//1、创建rpc服务器
	err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、监听端口
	listner, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
	}
	//3、退出时关闭服务
	defer listner.Close()
	for {
		fmt.Println("开始建立连接...")
		//4、建立连接
		conn, err3 := listner.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		//5、绑定服务
		rpc.ServeConn(conn)
	}

}

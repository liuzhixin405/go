package main

import (
	"context"
	"fmt"
	"mygreeter/proto/greeter"
	"net"

	"google.golang.org/grpc"
)

type Hello struct{}

func (s Hello) SayHello(ctx context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{Message: "我是服务端,Hello " + req.Name}, nil
}
func main() {
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	grpcServer.Serve(listener)
}

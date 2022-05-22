package main

import (
	"context"
	"log"
	"mygreeter/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := greeter.NewGreeterClient(grpcClient)
	res, err2 := client.SayHello(context.Background(), &greeter.HelloReq{Name: "张晓"})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Println("%v", res)
}

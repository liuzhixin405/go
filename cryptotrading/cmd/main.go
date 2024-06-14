package main

import (
	"cryptotrading/api"
	"cryptotrading/pb"
	"cryptotrading/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	go startGRPCServer()

	r := api.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTradeServiceServer(s, &service.TradeService{})

	log.Println("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

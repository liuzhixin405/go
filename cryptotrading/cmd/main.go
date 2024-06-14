package main

import (
	"cryptotrading/api"
	"cryptotrading/db"
	"cryptotrading/pb"
	"cryptotrading/service"
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	database := db.InitDB("root:my-secret-pw@tcp(127.0.0.1:3306)/cryptotrading")
	go startGRPCServer(database)

	r := api.SetupRouter(database)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run HTTP server: %v", err)
	}
}

func startGRPCServer(db *sql.DB) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tradeService := service.NewTradeService(db)
	s := grpc.NewServer()
	pb.RegisterTradeServiceServer(s, tradeService)

	log.Println("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

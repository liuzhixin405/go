package service

import (
	"context"
	"cryptotrading/pb"
	"database/sql"
	"fmt"
)

type TradeService struct {
	pb.UnimplementedTradeServiceServer
	DB *sql.DB
}

func NewTradeService(db *sql.DB) *TradeService {
	return &TradeService{DB: db}
}
func (s *TradeService) ExecuteTrade(ctx context.Context, req *pb.TradeRequest) (*pb.TradeResponse, error) {
	fmt.Printf("User %s wants to trade %f %s to %s\n", req.UserId, req.Amount, req.FromCurrency, req.ToCurrency)

	_, err := s.DB.Exec("INSERT INTO trades (user_id, from_currency, to_currency, amount) VALUES (?, ?, ?, ?)",
		req.UserId, req.FromCurrency, req.ToCurrency, req.Amount)
	if err != nil {
		return &pb.TradeResponse{
			Success: false,
			Message: "Failed to execute trade",
		}, err
	}
	return &pb.TradeResponse{
		Success: true,
		Message: "Trade executed successfully",
	}, nil
}

package service

import (
	"context"
	"cryptotrading/pb"
	"fmt"
)

type TradeService struct {
	pb.UnimplementedTradeServiceServer
}

func (s *TradeService) ExecuteTrade(ctx context.Context, req *pb.TradeRequest) (*pb.TradeResponse, error) {
	fmt.Printf("User %s wants to trade %f %s to %s\n", req.UserId, req.Amount, req.FromCurrency, req.ToCurrency)
	return &pb.TradeResponse{
		Success: true,
		Message: "Trade executed successfully",
	}, nil
}

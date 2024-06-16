package trade

import (
	"fmt"
	"mycrypto/order"
	"mycrypto/user"
)

type TradeService interface {
	ExecuteTrade(userID, orderID int, amount float64, direction string) (int, error)
}

type tradeService struct {
	tradeRepo   TradeRepository
	orderRepo   order.OrderRepository
	userService user.UserService
}

func NewTradeService(tradeRepo TradeRepository, orderRepo order.OrderRepository, userService user.UserService) TradeService {
	return &tradeService{
		tradeRepo:   tradeRepo,
		orderRepo:   orderRepo,
		userService: userService,
	}
}

func (s *tradeService) ExecuteTrade(userID, orderID int, amount float64, direction string) (int, error) {
	// check if user has enough balance
	user, err := s.userService.GetUserInfo(userID)
	if err != nil {
		return 0, err
	}
	if user.Balance < amount {
		return 0, fmt.Errorf("余额不足")
	}
	_, err = s.tradeRepo.SaveTrade(&Trade{
		UserID:    userID,
		OrderID:   orderID,
		Amount:    amount,
		Direction: Buy,
	})
	if err != nil {
		return 0, err
	}
	return 1, nil
}

package order

type OrderService interface {
	CreateOrder(userID int, fromCurrency string, toCurrency string, amount float64) (int, error)
}

type orderService struct {
	orderRepo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderService{
		orderRepo: repo,
	}
}

func (s *orderService) CreateOrder(userID int, fromCurrency string, toCurrency string, amount float64) (int, error) {
	order := &Order{
		UserID:       userID,
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Amount:       amount,
		Status:       "pending",
	}
	orderID, err := s.orderRepo.Create(order)
	if err != nil {
		return 0, err
	}
	return orderID, err
}

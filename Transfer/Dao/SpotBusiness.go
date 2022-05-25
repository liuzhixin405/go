package Dao

type SpotBusiness struct{}

func (s *SpotBusiness) GetCoinAvailableQuantity(token, coin string) float64 {

	return 1
}

func (s *SpotBusiness) ConfirmTransfer(orderId string, success bool) bool {
	return false
}

func (s *SpotBusiness) TransferAssets(token, orderId, coin string, amount float64, side int) bool {
	return false
}

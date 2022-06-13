package model

import "github.com/shopspring/decimal"

type Store interface {
	TransferAssets(token, orderId, coin string, amount decimal.Decimal, side TransferDirection) bool
	ConfirmTransfer(orderId string, success bool) bool
	GetCoinAvailableQuantity(customerId string, coin string) decimal.Decimal
}

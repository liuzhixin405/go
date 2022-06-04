package service

import (
	"transferasset/model"

	"github.com/shopspring/decimal"
)

func TransferAssets(token, orderId, coin string, amount decimal.Decimal, side model.TransferDirection) bool {
	return true
}

func ConfirmTransfer(orderId string, success bool) bool {

}

func GetCoinAvailableQuantity(customerId string, coin string) decimal.Decimal {

}

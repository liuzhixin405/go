package service

import (
	"transferasset/model"
	"transferasset/model/psql"

	"github.com/shopspring/decimal"
)

func TransferAssets(token, orderId, coin string, amount decimal.Decimal, side model.TransferDirection) bool {
	return psql.SharedStore().TransferAssets(token, orderId, coin, amount, side)
}

func ConfirmTransfer(orderId string, success bool) bool {
	return psql.SharedStore().ConfirmTransfer(orderId, success)
}

func GetCoinAvailableQuantity(customerId string, coin string) decimal.Decimal {
	//return decimal.NewFromInt32(123)
	return psql.SharedStore().GetCoinAvailableQuantity(customerId, coin)
}

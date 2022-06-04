package psql

import (
	"transferasset/model"

	"github.com/shopspring/decimal"
)

func (s *Store) GetCoinAvailableQuantity(customerId string, coin string) float64 {
	var asset model.Asset
	s.db.Where(&model.Asset{CustomerId: customerId, CoinId: coin, BusinessId: "1"}).First(&asset)
	if asset.AvaliableQuantity == 0 {
		return 0
	} else {
		return asset.AvaliableQuantity
	}
}

func (s *Store) ConfirmTransfer(orderId string, success bool) bool {
	if !success || orderId == "" {
		return false
	}
	// var assetTransferRecord model.AssetTransferRecord
	// if err := db.PsqlHelper.Where(&model.AssetTransferRecord{OrderId: orderId}).First(&assetTransferRecord).Error; err != nil {
	// 	return false
	// } else if assetTransferRecord.Status == model.Confirmed {
	// 	return true
	// } else if assetTransferRecord.Status == model.Undo {
	// 	return false
	// } else if assetTransferRecord.Status == model.Error {
	// 	return false
	// } else if assetTransferRecord.Status == model.Confirming {
	// 	return false
	// }
	// if model.AssetsWasteBookType(assetTransferRecord.Direction) == model.Out && !success {
	// 	assetTransferRecord.Status == model.Undo
	// } else {
	// 	assetTransferRecord.Status == model.Confirmed
	// }
	//资金记录状态更新
	//资产流水更新

	return true
}

func (s *Store) TransferAssets(token, orderId, coin string, amount decimal.Decimal, side model.TransferDirection) bool {
	//省略流程
	return true
}

/*
TransferAssets(token, orderId, coin string, amount decimal.Decimal, side TransferDirection) bool
	ConfirmTransfer(orderId string, success bool) bool
	GetCoinAvailableQuantity(customerId string, coin string) decimal.Decimal*/

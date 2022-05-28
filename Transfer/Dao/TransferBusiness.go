package Dao

import (
	"transferasset/db"
	"transferasset/model"
)

type SpotBusiness struct{}

func (s *SpotBusiness) GetCoinAvailableQuantity(customerId string, coin string) float64 {
	var asset model.Asset
	db.PsqlHelper.Where(&model.Asset{CustomerId: customerId, CoinId: coin, BusinessId: "1"}).First(&asset)
	if asset.AvaliableQuantity == 0 {
		return 0
	} else {
		return asset.AvaliableQuantity
	}
}

func (s *SpotBusiness) ConfirmTransfer(orderId string, success bool) bool {
	if !success || orderId == "" {
		return false
	}
	var assetTransferRecord model.AssetTransferRecord
	if err := db.PsqlHelper.Where(&model.AssetTransferRecord{OrderId: orderId}).First(&assetTransferRecord).Error; err != nil {
		return false
	} else if assetTransferRecord.Status == model.Confirmed {
		return true
	} else if assetTransferRecord.Status == model.Undo {
		return false
	} else if assetTransferRecord.Status == model.Error {
		return false
	} else if assetTransferRecord.Status == model.Confirming {
		return false
	}
	if assetTransferRecord.Direction == model.Out && !success {
		assetTransferRecord.Status == model.Undo
	} else {
		assetTransferRecord.Status == model.Confirmed
	}
	//资金记录状态更新
	//资产流水更新

	return true
}

func (s *SpotBusiness) TransferAssets(token, orderId, coin string, amount float64, side int) bool {
	//省略流程
	return true
}

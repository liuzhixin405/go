package model

import "github.com/shopspring/decimal"

//import "gorm.io/gorm"
type Asset struct {
	Dto
	CustomerId        string
	BusinessId        string
	CoinId            string
	AvaliableQuantity decimal.Decimal //可用
	FrozenQuantity    decimal.Decimal //占用
	Include           decimal.Decimal //划入
	Drawout           decimal.Decimal //划出
}

type AssetTransferRecord struct {
	Dto
	CstomerId  string
	CoiId      string
	Direction  TransferDirection
	Amount     decimal.Decimal
	Status     TransferStatus
	BusinessId string
	OrderId    string
}

type AssetWasteBook struct {
	Dto
	CustomerId                string
	BusinessId                string
	CoinId                    string
	Remarks                   string
	AssociatedOrderId         string
	OriginalAvailableQuantity decimal.Decimal     //原始可用
	ChangeQuantity            decimal.Decimal     //更改数量
	AvailableQuantity         decimal.Decimal     //可用数量
	OriginalFrozenQuantity    decimal.Decimal     //原始冻结
	ChangeFrozenQuantity      decimal.Decimal     //更改冻结
	FrozenQuantity            decimal.Decimal     //冻结
	AssetsWasteBookType       AssetsWasteBookType //流水类型
}

type TransferDirection int

const (
	ToThird TransferDirection = iota
	InThird
)

type TransferStatus int

const (
	//确认中
	Confirming TransferStatus = iota
	//已确认
	Confirmed
	//取消
	Undo
	//错误
	Error
)

type AssetsWasteBookType int

const (
	In  AssetsWasteBookType = iota //转入
	Out                            //转出
)

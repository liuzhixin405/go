package model

type TransferDirection int

const (
	FuturesToSpot TransferDirection = iota
	SpotToFutures
	FuturesToFiat
	FiatToFutures
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

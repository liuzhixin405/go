package model

type AssetWasteBook struct {
	Dto
	CustomerId                string
	BusinessId                string
	CoinId                    string
	Remarks                   string
	AssociatedOrderId         string
	OriginalAvailableQuantity float64 //原始可用
	ChangeQuantity            float64 //更改数量
	AvailableQuantity         float64 //可用数量
	OriginalFrozenQuantity    float64 //原始冻结
	ChangeFrozenQuantity      float64 //更改冻结
	FrozenQuantity            float64 //冻结
	AssetsWasteBookType       int     //流水类型
}

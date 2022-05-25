package model

type Asset struct {
	Dto
	CustomerId        string
	BusinessId        string
	CoinId            string
	AvaliableQuantity float64 //可用
	FrozenQuantity    float64 //占用
	Include           float64 //划入
	Drawout           float64 //划出
}

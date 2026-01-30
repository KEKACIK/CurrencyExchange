package model

type ExchangeRate struct {
	Id               int
	BaseCurrencyId   int
	TargetCurrencyId int
	Rate             float32
}

func (ExchangeRate) TableName() string {
	return "exchangerates"
}

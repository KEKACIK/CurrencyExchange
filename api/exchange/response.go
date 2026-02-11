package exchange

import "CurrencyExchange/api/currency"

type Response struct {
	BaseCurrency    currency.Response `json:"baseCurrency"`
	TargetCurrency  currency.Response `json:"targetCurrency"`
	Rate            float32           `json:"rate"`
	Amount          float32           `json:"amount"`
	ConvertedAmount float32           `json:"convertedAmount"`
}

package exchangerate

import "CurrencyExchange/api/currency"

type Response struct {
	Id             int               `json:"id"`
	BaseCurrency   currency.Response `json:"baseCurrency"`
	TargetCurrency currency.Response `json:"targetCurrency"`
	Rate           float64           `json:"rate"`
}

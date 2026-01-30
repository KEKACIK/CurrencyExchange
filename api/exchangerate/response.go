package exchangerate

import "CurrencyExchange/api/currency"

type Response struct {
	id             int
	baseCurrency   currency.Response
	targetCurrency currency.Response
	rate           float32
}

package exchangerate

import "CurrencyExchange/api/currency"

type Response struct {
	baseCurrency    currency.Response
	targetCurrency  currency.Response
	rate            float32
	amount          float32
	convertedAmount float32
}

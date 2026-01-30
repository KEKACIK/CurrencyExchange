package main

import (
	"CurrencyExchange/api/currency"
	"CurrencyExchange/api/exchangerate"
	"fmt"
	"net/http"
)

func main() {
	currency.RouterInit()
	exchangerate.RouterInit()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

package exchangerate

import "net/http"

func RouterInit() {
	handler := NewHandler()

	http.HandleFunc("POST /exchangeRates", handler.Create)
	http.HandleFunc("GET /exchangeRate/{splitCode}", handler.Get)
	http.HandleFunc("GET /exchangeRates", handler.GetList)
	http.HandleFunc("PATCH /exchangeRates", handler.Update)
}

package currency

import "net/http"

func RouterInit() {
	handler := NewHandler()

	http.HandleFunc("POST /currencies", handler.Create)
	http.HandleFunc("GET /currency/{code}", handler.Get)
	http.HandleFunc("GET /currencies", handler.GetList)
}

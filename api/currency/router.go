package currency

import "net/http"

func RouterInit() {
	handler := NewHandler()

	http.HandleFunc("GET /currencies", handler.GetList)
	http.HandleFunc("POST /currencies", handler.Create)
	http.HandleFunc("GET /currency/{code}", handler.Get)
}

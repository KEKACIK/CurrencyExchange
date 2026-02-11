package exchangerate

import (
	"database/sql"
	"net/http"
)

func RouterInit(db *sql.DB) {
	handler := NewHandler(db)

	http.HandleFunc("POST /exchangeRates", handler.Create)
	http.HandleFunc("GET /exchangeRate/{splitCode}", handler.Get)
	http.HandleFunc("GET /exchangeRates", handler.GetList)
	http.HandleFunc("PATCH /exchangeRates", handler.Update)
}

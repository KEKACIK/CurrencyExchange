package currency

import (
	"database/sql"
	"net/http"
)

func RouterInit(db *sql.DB) {
	handler := NewHandler(db)

	http.HandleFunc("POST /currencies", handler.Create)
	http.HandleFunc("GET /currency/{code}", handler.Get)
	http.HandleFunc("GET /currencies", handler.GetList)
}

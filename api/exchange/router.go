package exchange

import (
	"database/sql"
	"net/http"
)

func RouterInit(db *sql.DB) {
	handler := NewHandler(db)

	http.HandleFunc("GET /exchange", handler.Get)
}

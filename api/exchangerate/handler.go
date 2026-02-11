package exchangerate

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	splitCode := r.PathValue("splitCode")
	baseCode, targetCode := splitCode[:3], splitCode[3:]
	fmt.Println(baseCode)
	fmt.Println(targetCode)
}

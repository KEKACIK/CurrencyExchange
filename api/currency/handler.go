package currency

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	fmt.Println(code)
}

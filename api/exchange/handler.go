package exchangerate

import (
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

}

package exchangerate

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
	splitCode := r.PathValue("splitCode")
	baseCode, targetCode := splitCode[:3], splitCode[3:]
	fmt.Println(baseCode)
	fmt.Println(targetCode)
}

package currency

import (
	"CurrencyExchange/internal/repository"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	db   *sql.DB
	repo repository.Currency
}

func NewHandler(db *sql.DB) *Handler {
	repo := repository.NewCurrency(db)

	return &Handler{
		db:   db,
		repo: repo,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	code := strings.ToUpper(r.FormValue("code"))
	if err := CodeValidation(code); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.repo.GetByCode(code)
	if err != nil && err != repository.ErrCurrencyNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err == nil {
		http.Error(w, "Currency already exists", http.StatusConflict)
		return
	}

	name := r.FormValue("name")
	if err := NameValidation(name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sign := r.FormValue("sign")
	if err := SignValidation(sign); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currency, err := h.repo.Create(code, name, sign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Response{
		Id:   currency.Id,
		Name: currency.FullName,
		Code: currency.Code,
		Sign: currency.Sign,
	})
}

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := []Response{}
	for _, v := range *currencies {
		result = append(result, Response{
			Id:   v.Id,
			Name: v.FullName,
			Code: v.Code,
			Sign: v.Sign,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	code := strings.ToUpper(r.PathValue("code"))
	if err := CodeValidation(code); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currency, err := h.repo.GetByCode(code)
	if err != nil {
		if err == repository.ErrCurrencyNotFound {
			http.Error(w, "Currency not found", http.StatusNotFound)
		}
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Response{
		Id:   currency.Id,
		Name: currency.FullName,
		Code: currency.Code,
		Sign: currency.Sign,
	})

}

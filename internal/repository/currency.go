package repository

import (
	"CurrencyExchange/internal/model"
	"database/sql"
	"errors"
)

type Currency interface {
	Create(code, fullName, sign string) (*model.Currency, error)
	GetAll() (*[]model.Currency, error)
	GetByCode(code string) (*model.Currency, error)
	Update(code, fullName, sign string) (*model.Currency, error)
}

type currency struct {
	db *sql.DB
}

var (
	ErrCurrencyNotFound = errors.New("invalid: currency not found")
)

func NewCurrency(db *sql.DB) Currency {
	return &currency{db: db}
}

func (c *currency) Create(code, fullName, sign string) (*model.Currency, error) {
	_, err := c.db.Exec("INSERT INTO currencies(code, fullname, sign) VALUES(?, ?, ?)", code, fullName, sign)
	if err != nil {
		return nil, err
	}
	return c.GetByCode(code)
}

func (c *currency) GetAll() (*[]model.Currency, error) {
	rows, err := c.db.Query("SELECT * FROM currencies")
	if err != nil {
		return nil, err
	}

	currencies := []model.Currency{}
	for rows.Next() {
		currency := model.Currency{}
		err = rows.Scan(&currency.Id, &currency.Code, &currency.FullName, &currency.Sign)
		currencies = append(currencies, currency)
	}

	return &currencies, nil
}

func (c *currency) GetByCode(code string) (*model.Currency, error) {
	rows, err := c.db.Query("SELECT * FROM currencies WHERE code=?", code)
	if err != nil {
		return nil, err
	}

	var currency *model.Currency
	for rows.Next() {
		currency = &model.Currency{}
		err = rows.Scan(&currency.Id, &currency.Code, &currency.FullName, &currency.Sign)
	}
	if currency == nil {
		return nil, ErrCurrencyNotFound
	}

	return currency, nil
}

func (c *currency) Update(code, fullName, sign string) (*model.Currency, error) {
	_, err := c.db.Exec("UPDATE currencies SET fullname = ?, sign = ? WHERE code=?", fullName, sign, code)
	if err != nil {
		return nil, err
	}

	return c.GetByCode(code)
}

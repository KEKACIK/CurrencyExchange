package repository

import (
	"CurrencyExchange/internal/model"
	"database/sql"
)

type ExchangeRate interface {
	Create(BaseCurrencyId, TargetCurrencyId int, rate float32) (*model.ExchangeRate, error)
	GetByCurrenciesId(baseCurrencyId, targetCurrencyId int) (*model.ExchangeRate, error)
}

type exchangerate struct {
	db *sql.DB
}

func NewExchangeRate(db *sql.DB) ExchangeRate {
	return &exchangerate{db: db}
}

func (ec *exchangerate) Create(baseCurrencyId, targetCurrencyId int, rate float32) (*model.ExchangeRate, error) {
	_, err := ec.db.Exec("INSERT INTO currencies(code, fullname, sign) VALUES(?, ?, ?)", baseCurrencyId, targetCurrencyId, rate)
	if err != nil {
		return nil, err
	}
	return ec.GetByCurrenciesId(baseCurrencyId, targetCurrencyId)
}

func (ec *exchangerate) GetByCurrenciesId(baseCurrencyId, targetCurrencyId int) (*model.ExchangeRate, error) {
	rows, err := ec.db.Query("SELECT * FROM %s WHERE basecurrencyid=? AMD targetcurrencyid", baseCurrencyId, targetCurrencyId)
	if err != nil {
		return nil, err
	}

	exchangeRate := model.ExchangeRate{}
	for rows.Next() {
		err = rows.Scan(&exchangeRate.Id, &exchangeRate.BaseCurrencyId, &exchangeRate.TargetCurrencyId, &exchangeRate.Rate)
	}

	return &exchangeRate, nil
}

// func (ec *exchangerate) Update(baseCurrencyId, targetCurrencyId int, rate float32) (*model.ExchangeRate, error) {
// 	_, err := ec.db.Exec("UPDATE %s SET fullname = ?, sign = ? WHERE basecurrencyid=? AND targetCurrencyId=?", fullName, sign, code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ec.GetByCode(code)
// }

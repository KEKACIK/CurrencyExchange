package database

import (
	"CurrencyExchange/internal/model"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type mathFunc func(*sql.DB)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "assets/database.db")
	if err != nil {
		return nil, err
	}
	initFunctions := []func(*sql.DB) error{
		model.NewCurrency,
		model.NewExchangeRate,
	}

	for _, fn := range initFunctions {
		err = fn(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

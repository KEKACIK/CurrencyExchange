package model

import "database/sql"

type ExchangeRate struct {
	Id               int
	BaseCurrencyId   int
	TargetCurrencyId int
	Rate             float32
}

func NewExchangeRate(db *sql.DB) error {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS exchangerates (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        basecurrencyid INTEGER,
        targetcurrencyid INTEGER,
        rate DECIMAL(8)
    );
    `
	_, err := db.Exec(sqlStmt)

	return err
}

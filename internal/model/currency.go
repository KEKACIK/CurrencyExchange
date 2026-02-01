package model

import (
	"database/sql"
)

type Currency struct {
	Id       int
	Code     string
	FullName string
	Sign     string
}

func NewCurrency(db *sql.DB) error {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS currencies (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        code VARCHAR(8) UNIQUE,
        fullname VARCHAR(256),
        sign VARCHAR(4)
    );
    `
	_, err := db.Exec(sqlStmt)

	return err
}

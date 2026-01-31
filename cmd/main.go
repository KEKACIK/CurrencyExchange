package main

import (
	"CurrencyExchange/api/currency"
	"CurrencyExchange/api/exchangerate"
	"CurrencyExchange/pkg/database"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Database init...")
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Api routers init...")
	currency.RouterInit()
	exchangerate.RouterInit()

	log.Println("Api server started")
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

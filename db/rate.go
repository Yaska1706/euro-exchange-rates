package db

import (
	"database/sql"
	"log"
)

type CurrencyRate struct {
	Currency string
	Rate     int
	Date     string
}

func Create(db *sql.DB, currency, rate, date string) {
	_, err := db.Exec(`INSERT INTO "public"."rates"("currency","rate","date") VALUES($1,$2,$3)`, currency, rate, date)
	if err != nil {
		log.Print(err)
	}

}

func GetLatest(db *sql.DB, date string) []CurrencyRate {
	currencyrates := []CurrencyRate{}
	rows, err := db.Query(`SELECT "currency", "rate" FROM "public"."rates" WHERE "date" = $1`, date)
	if err != nil {
		log.Print("💀 error: ", err)
	}
	for rows.Next() {
		currencyrate := CurrencyRate{}
		rows.Scan(&currencyrate.Currency, &currencyrate.Rate)
		currencyrates = append(currencyrates, currencyrate)
	}
	return currencyrates
}

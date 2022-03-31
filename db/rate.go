package db

import (
	"database/sql"
	"log"
)

type CurrencyRate struct {
	Currency string
	Rate     string
	Date     string
}

func Create(db *sql.DB, exrate CurrencyRate) {
	_, err := db.Exec(`INSERT INTO "public"."rates"("currency","rate","date") VALUES($1,$2,$3)`, exrate.Currency, exrate.Rate, exrate.Date)
	if err != nil {
		log.Print(err)
	}

}

func GetByDate(db *sql.DB, date string) []CurrencyRate {
	currencyrates := []CurrencyRate{}
	rows, err := db.Query(`SELECT "currency", "rate" FROM "public"."rates" WHERE "date" = $1`, date)
	if err != nil {
		log.Print("error: ", err)
	}
	for rows.Next() {
		currencyrate := CurrencyRate{}
		if err := rows.Scan(&currencyrate.Currency, &currencyrate.Rate); err != nil {
			log.Print("Error:", err)

		}
		currencyrates = append(currencyrates, currencyrate)
	}
	return currencyrates
}

func GetAllRates(db *sql.DB) []CurrencyRate {
	currencyrates := []CurrencyRate{}
	rows, err := db.Query(`SELECT * FROM "public"."rates"`)
	if err != nil {
		log.Print("error: ", err)
	}
	for rows.Next() {
		currencyrate := CurrencyRate{}
		var id int
		if err := rows.Scan(&id, &currencyrate.Currency, &currencyrate.Date, &currencyrate.Rate); err != nil {
			log.Print("Error:", err)

		}
		currencyrates = append(currencyrates, currencyrate)
	}
	return currencyrates
}

func GetByCurrency(db *sql.DB, currency string) []string {
	rates := []string{}
	rows, err := db.Query(`SELECT "rate" FROM "public"."rates" WHERE "currency" = $1`, currency)
	if err != nil {
		log.Print("error: ", err)
	}
	for rows.Next() {
		currencyrate := CurrencyRate{}

		if err := rows.Scan(&currencyrate.Rate); err != nil {
			log.Print("Error:", err)

		}
		rates = append(rates, currencyrate.Rate)
	}
	return rates
}

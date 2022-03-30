package api

import (
	"time"

	"github.com/yaska1706/rakuten-interview/db"
)

type ExchangeRate struct {
	Date     string `json:"date"`
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
}

func getlatestdate() string {
	t := time.Now()
	currentDate := t.Format("2006-03-30")

	return currentDate
}

func returnlatestrates() []ExchangeRate {
	DB := db.DBConnection()
	server := &server{
		db: DB,
	}
	var exchangeRates []ExchangeRate

	date := getlatestdate()
	latestcurrencyrates := db.GetByDate(server.db, date)
	for _, latestcurrencyrate := range latestcurrencyrates {
		exchangeRate := ExchangeRate{
			Currency: latestcurrencyrate.Currency,
			Rate:     latestcurrencyrate.Rate,
		}
		exchangeRates = append(exchangeRates, exchangeRate)
	}
	return exchangeRates
}

func returnratesperdate(date string) []ExchangeRate {
	DB := db.DBConnection()
	server := &server{
		db: DB,
	}
	var exchangeRates []ExchangeRate
	latestcurrencyrates := db.GetByDate(server.db, date)
	for _, latestcurrencyrate := range latestcurrencyrates {
		exchangeRate := ExchangeRate{
			Currency: latestcurrencyrate.Currency,
			Rate:     latestcurrencyrate.Rate,
		}
		exchangeRates = append(exchangeRates, exchangeRate)
	}
	return exchangeRates
}

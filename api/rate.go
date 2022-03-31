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
	t := time.Now().Local()
	currentDate := t.Format("2006-01-02")

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

func returnAllRates() []ExchangeRate {
	DB := db.DBConnection()
	server := &server{
		db: DB,
	}
	var exchangeRates []ExchangeRate
	allrates := db.GetAllRates(server.db)
	for _, allrate := range allrates {
		exchangeRate := ExchangeRate{
			Currency: allrate.Currency,
			Rate:     allrate.Rate,
		}
		exchangeRates = append(exchangeRates, exchangeRate)
	}
	return exchangeRates
}

// func findMin(exchangerates []ExchangeRate) string {
// 	exchangerates = returnAllRates()
// 	for _, exchangerate := range exchangerates {
// 		exchangerate.Currency
// 	}
// 	return ""
// }

func MinMax(array []string) (string, string) {
	var max string = array[0]
	var min string = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

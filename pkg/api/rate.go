package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yaska1706/rakuten-interview/pkg/db"
)

type ExchangeRate struct {
	Date     string `json:"date"`
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
}

func getlatestdate() string {
	t := time.Now().UTC()
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
	if len(latestcurrencyrates) == 0 {
		t := time.Now().UTC()
		newdate := t.AddDate(0, 0, -1)
		date = newdate.Format("2006-01-02")
		latestcurrencyrates = db.GetByDate(server.db, date)
	}
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

func returnratespercurrency(currency string) []string {
	DB := db.DBConnection()
	server := &server{
		db: DB,
	}
	rates := db.GetByCurrency(server.db, currency)

	return rates
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

func storecurrencies() []string {
	var currencies []string

	exchangerates := returnAllRates()
	for _, exchangerate := range exchangerates {
		if contains(currencies, exchangerate.Currency) {
			continue
		}
		currencies = append(currencies, exchangerate.Currency)
	}
	return currencies
}

type CurrencyPerRate struct {
	Name  string
	Rates []string
}

func storeratespercurrency() []CurrencyPerRate {
	currencyperrates := []CurrencyPerRate{}
	currencies := storecurrencies()

	for _, currency := range currencies {

		rates := returnratespercurrency(currency)
		currencyperrate := CurrencyPerRate{
			Name:  currency,
			Rates: rates,
		}
		currencyperrates = append(currencyperrates, currencyperrate)

	}

	return currencyperrates
}

type AnalysisValues struct {
	Currency string
	Values   map[string]string
}

type AnalyzeRates map[string]interface{}

func getMinMaxrates() AnalyzeRates {

	analyzerates := AnalyzeRates{}
	currencyperates := storeratespercurrency()
	values := map[string]string{}
	for _, currencyperate := range currencyperates {

		Min, Max := MinMax(currencyperate.Rates)
		Avg := Average(currencyperate.Rates)

		values["min"] = Min
		values["max"] = Max
		values["Avg"] = Avg
		analyzerates[currencyperate.Name] = values

	}
	return analyzerates
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func MinMax(array []string) (string, string) {
	arraytofloat := []float64{}
	for _, value := range array {
		if n, err := strconv.ParseFloat(value, 64); err == nil {
			arraytofloat = append(arraytofloat, n)
		}
	}

	var max float64 = arraytofloat[0]
	var min float64 = arraytofloat[0]

	for _, value := range arraytofloat {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	maxvalue := fmt.Sprint(max)
	minvalue := fmt.Sprint(min)
	return minvalue, maxvalue
}

func Average(array []string) string {
	arraytofloat := []float64{}
	for _, value := range array {
		if n, err := strconv.ParseFloat(value, 64); err == nil {
			arraytofloat = append(arraytofloat, n)
		}
	}
	var sum float64

	for _, value := range arraytofloat {
		sum += value
	}
	n := len(arraytofloat)
	avg := fmt.Sprint(sum / float64(n))

	return avg
}

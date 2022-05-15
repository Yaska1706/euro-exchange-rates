package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/antchfx/xmlquery"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yaska1706/rakuten-interview/pkg/api"
	"github.com/yaska1706/rakuten-interview/pkg/db"
)

func initialize() error {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	DB := db.DBConnection()
	if err := db.SeedDB(DB); err != nil {
		return fmt.Errorf("Seeding DB : %w", err)
	}
	if err := SaveXMLToDB(DB); err != nil {
		return fmt.Errorf("savetoxml: %w", err)
	}

	return nil
}
func main() {
	if err := initialize(); err != nil {
		log.Fatal("Initialize : %w", err)
		os.Exit(1)
	}
	router := mux.NewRouter()
	serve := api.NewServer(router)
	serve.Run()

}

func queryxmldata() ([]db.CurrencyRate, error) {
	doc, err := xmlquery.LoadURL("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		return nil, fmt.Errorf("xmlquery : %w", err)
	}
	currencyrates := []db.CurrencyRate{}
	var datetime string

	lists, _ := xmlquery.QueryAll(doc, "//Cube//Cube")

	var currencyrate db.CurrencyRate
	for _, list := range lists {

		if list.SelectAttr("time") == "" {
			continue
		}
		datetime = list.SelectAttr("time")
		for _, value := range list.SelectElements("//Cube") {
			if value.SelectAttr("currency") == "" || value.SelectAttr("currency") == "" {
				continue
			}

			currencyrate.Date = datetime
			currencyrate.Currency = value.SelectAttr("currency")
			currencyrate.Rate = value.SelectAttr("rate")

			currencyrates = append(currencyrates, currencyrate)
		}
	}
	return currencyrates, nil
}

func SaveXMLToDB(DB *sql.DB) error {
	currencyrates, err := queryxmldata()
	if err != nil {
		return fmt.Errorf("queryxml : %w", err)
	}
	for _, currencyrate := range currencyrates {
		db.Create(DB, currencyrate)
	}

	return nil
}

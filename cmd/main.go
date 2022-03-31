package main

import (
	"database/sql"
	"log"

	"github.com/antchfx/xmlquery"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yaska1706/rakuten-interview/pkg/api"
	"github.com/yaska1706/rakuten-interview/pkg/db"
)

func init() {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	DB := db.DBConnection()
	db.SeedDB(DB)
	SaveXMLToDB(DB)

}
func main() {
	router := mux.NewRouter()
	serve := api.NewServer(router)
	serve.Run()

}

func queryxmldata() []db.CurrencyRate {
	doc, err := xmlquery.LoadURL("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Print(err)
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
	return currencyrates
}

func SaveXMLToDB(DB *sql.DB) {
	currencyrates := queryxmldata()
	for _, currencyrate := range currencyrates {
		db.Create(DB, currencyrate)
	}
}

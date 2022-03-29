package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/antchfx/xmlquery"
	"github.com/yaska1706/rakuten-interview/db"
)

func main() {

	DB := db.DBConnection()

	db.SeedDB(DB)
	usingxmlquery(DB)
}

type CurrencyTime struct {
	Currency string
	Rate     string
}
type CurrencyValues struct {
	Time         string
	CurrencyTime []CurrencyTime
}

func usingxmlquery(DB *sql.DB) {
	doc, err := xmlquery.LoadURL("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Print(err)
	}

	lists, _ := xmlquery.QueryAll(doc, "//Cube//Cube")
	var datetime string
	for _, list := range lists {

		if list.SelectAttr("time") == "" {
			continue
		}
		fmt.Printf("Time: %s\n", list.SelectAttr("time"))
		datetime = list.SelectAttr("time")
		for _, value := range list.SelectElements("//Cube") {
			if value.SelectAttr("currency") == "" || value.SelectAttr("currency") == "" {
				continue
			}
			fmt.Printf("Currency: %s Rate %s\n", value.SelectAttr("currency"), value.SelectAttr("rate"))
			db.Create(DB, value.SelectAttr("currency"), value.SelectAttr("rate"), datetime)
		}
	}
}

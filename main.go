package main

import (
	"fmt"
	"log"

	"github.com/antchfx/xmlquery"
	"github.com/yaska1706/rakuten-interview/db"
)

func main() {

	DB := db.DBConnection()

	db.SeedDB(DB)
	usingxmlquery()
}

type CurrencyTime struct {
	Currency string
	Rate     string
}
type CurrencyValues struct {
	Time         string
	CurrencyTime []CurrencyTime
}

func usingxmlquery() {
	doc, err := xmlquery.LoadURL("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Print(err)
	}

	lists, _ := xmlquery.QueryAll(doc, "//Cube//Cube")
	for _, list := range lists {
		if list.SelectAttr("time") == "" {
			continue
		}
		fmt.Printf("Time: %s\n", list.SelectAttr("time"))
		for _, value := range list.SelectElements("//Cube") {
			if value.SelectAttr("currency") == "" || value.SelectAttr("currency") == "" {
				continue
			}
			fmt.Printf("Currency: %s Rate %s\n", value.SelectAttr("currency"), value.SelectAttr("rate"))
		}
	}
}

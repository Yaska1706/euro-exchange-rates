package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "development"
	password = "develop"
	dbname   = "event_demo"
)

func DBConnection() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func SeedDB(db *sql.DB) error {
	log.Print("💾 Seeding database with table...")

	_, err := db.Exec(`
		DROP TABLE IF EXISTS rates;
		CREATE TABLE IF NOT EXISTS rates (
			"id"      SERIAL PRIMARY KEY,
			"currency"    varchar(50) NOT NULL,
			"date"   varchar(50) NOT NULL,
			"rate"    varchar(50) NOT NULL
		)
	`)

	if err != nil {
		log.Print("query error: ", err)
	}

	return err
}

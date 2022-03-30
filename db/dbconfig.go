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

func deletestableifexists(db *sql.DB) {
	log.Print("Deleting existing table...")
	_, err := db.Exec(`DROP TABLE rates`)
	if err != nil {
		log.Print("error:", err)
		return
	}
}

func SeedDB(db *sql.DB) error {
	deletestableifexists(db)
	log.Print("💾 Seeding database with table...")

	_, err := db.Exec(`
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

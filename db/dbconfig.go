package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DBConnection() *sql.DB {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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

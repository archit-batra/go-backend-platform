package testutil

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func SetupTestDB() *sql.DB {

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=wallet sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

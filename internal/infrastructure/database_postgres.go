package infrastructure

import (
	"database/sql"
	"log"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=user sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	return db, nil
}

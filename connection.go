package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func connection() *sql.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE room (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			mobile VARCHAR(20) NOT NULL);
	`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	return db
}

package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	connStr := "postgres://postgres:secret@localhost:5432/myapp?sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Connected to PostgreSQL ✅")
}

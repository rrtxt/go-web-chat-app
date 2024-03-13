package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() {
	if db != nil {
		return
	}

	connStr := os.Getenv("DB_URI")
	if connStr == "" {
		log.Fatal("DB_URI must be set on environment variable")
	}

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db = client
}

func GetDB() *sql.DB {
	if db == nil {
		ConnectDB()
	}

	return db
}
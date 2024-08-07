package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func LoadDB() {
	var err error
	DB, err = InitDB("edu.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
}

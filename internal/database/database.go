package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Opens database and create database tables
func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		return fmt.Errorf("Unable to open database: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("Fail to ping database: %w", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS
   users (userID TEXT PRIMARY KEY, name TEXT)`)

	return err
}

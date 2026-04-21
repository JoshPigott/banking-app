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
		return fmt.Errorf("Unable to open database: %w\n", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("Fail to ping database: %w\n", err)
	}

	// Users
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS
	 users (userID TEXT PRIMARY KEY,
	 username TEXT UNIQUE, hashedPassword TEXT)`)

	// Sessions
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS
	sessions (id TEXT PRIMARY KEY NOT NULL,
	userID TEXT, expiryTime INTEGER)`)

	// Bank accounts
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS
	everydayAccount (userID TEXT FORGIN KEY,
	balanceCents INTEGER)`)

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS
	saverAccount (userID TEXT FORGIN KEY,
	balanceCents INTEGER)`)

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS
	kiwiSaverAccount (userID TEXT FORGIN KEY,
	balanceCents INTEGER)`)

	return err
}

package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return err
	}

	// Ensure table exists
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS resources (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
	);`)
	if err != nil {
		return err
	}

	log.Println("Database connection successfully initialized.")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
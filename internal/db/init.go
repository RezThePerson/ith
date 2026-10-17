package db

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

//go:embed schema.sql
var schema string

func Open() error {
	// Opening database
	var err error
	if DB, err = sql.Open("sqlite", "frosting.sqlite3"); err != nil {
		return fmt.Errorf("opening database: %w", err)
	}

	// Verifying database connection
	if err := DB.Ping(); err != nil {
		DB.Close()
		return fmt.Errorf("verifying database connection: %w", err)
	}

	// Executing database schema
	if _, err := DB.Exec(schema); err != nil {
		DB.Close()
		return fmt.Errorf("executing database schema: %w", err)
	}

	return nil
}

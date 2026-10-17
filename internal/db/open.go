package db

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

//go:embed schema.sql
var schema string

func Open() (err error) {
	// Getting database URL
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Opening database
	if DB, err = sql.Open("postgres", url); err != nil {
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

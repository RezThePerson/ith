package main

import (
	"log/slog"
	"os"

	"github.com/reztheperson/ith/internal/api"
	"github.com/reztheperson/ith/internal/db"
)

func main() {
	// Opening DB
	if err := db.Open(); err != nil {
		slog.Error("Failed to open DB", "err", err)
		os.Exit(1)
	}

	
	
	// Starting Api
	if err := api.Start(); err != nil {
		slog.Error("Failed to start api", "err", err)
		os.Exit(1)
	}
}

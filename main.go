package main

import (
	"log/slog"
	"os"

	"github.com/reztheperson/ith/internal/db"
)

func main() {
	// Opening DB
	if err := db.Open(); err != nil {
		slog.Error("Failed to open DB", "err", err)
		os.Exit(1)
	}
	slog.Info("DB Opened")

	// Starting Api

	slog.Info("Api Started")

}

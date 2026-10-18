package main

import (
	"log/slog"
	"os"

	"github.com/reztheperson/ith/internal/api"
)

func main() {
	// Opening DB
	/*
		if err := db.Open(); err != nil {
			slog.Error("Failed to open DB", "err", err)
			os.Exit(1)
		}
	*/

	// Starting API
	if err := api.Start(); err != nil {
		slog.Error("Failed to start API", "err", err)
		os.Exit(1)
	}
}

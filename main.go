package main

import (
	"log/slog"
	"os"

	"github.com/reztheperson/ith/internal/http"
)

func main() {
	// Opening DB
	/*
		if err := db.Open(); err != nil {
			slog.Error("Failed to open DB", "err", err)
			os.Exit(1)
		}
	*/

	// Starting http
	if err := http.Start(); err != nil {
		slog.Error("Failed to start http", "err", err)
		os.Exit(1)
	}
}

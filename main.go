package main

import (
	"log/slog"
	"os"
)

func main() {
	// Opening DB
	if err := db.Init(); err != nil {
		slog.Error("Failed to initialize DB", "err", err)
		os.Exit(1)
	}
	slog.Info("DB Initialized")

	// Starting Web Server
	slog.Info("Starting Web Server")

	// Starting Deployment Rounds
	slog.Info("Starting Deployment Rounds")
	rounds.Start()
}

package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/reztheperson/ith/internal/web"
)

func Start() error {
	mux := http.NewServeMux()

	// Set up routes
	if err := setupAPIRoutes(mux); err != nil {
		return fmt.Errorf("initializing api routes: %w", err)
	}
	if err := web.InitTemplates(); err != nil {
		return fmt.Errorf("initializing templates: %w", err)
	}
	if err := setupAppRoutes(mux); err != nil {
		return fmt.Errorf("setting up app routes: %w", err)
	}

	// Serving routes
	slog.Info("App successfully started", "port", ":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return fmt.Errorf("serving api routes: %w", err)
	}

	return nil
}

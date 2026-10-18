package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reztheperson/ith/internal/web"
)

func Start() (err error) {
	r := mux.NewRouter()

	// Set up routes
	if setupApiRoutes(r); err != nil {
		return fmt.Errorf("Initializing api routes: %w", err)
	}
	if web.InitTemplates(); err != nil {
		return fmt.Errorf("Initializing templates: %w", err)
	}
	if setupAppRoutes(r); err != nil {
		return fmt.Errorf("Setting up app routes: %w", err)
	}

	// Serving routes
	slog.Info("App successfully started", "port", ":8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("serving api routes: %w", err)
	}

	return nil
}

package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() error {
	r := mux.NewRouter()

	// Set up routes
	setupAPIRoutes(r)
	setupPageRoutes(r)

	// Serving routes
	slog.Info("App successfully started", "port", ":8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("serving http routes: %w", err)
	}

	return nil
}

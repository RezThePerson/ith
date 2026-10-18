package api

import (
	"fmt"
	"net/http"

	"github.com/reztheperson/ith/internal/web"
)

func setupAPIRoutes(mux *http.ServeMux) error {
	// API routes
	return nil
}

func setupAppRoutes(mux *http.ServeMux) error {
	// Static files
	if err := web.SetupStatic(mux); err != nil {
		return fmt.Errorf("setting up static routes: %w", err)
	}

	// Pages
	mux.HandleFunc("GET /{$}", HomeHandler)

	return nil
}

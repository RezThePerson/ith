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
	setupAPIRoutes(mux)
	mux.Handle("GET /static/", http.FileServer(http.FS(web.StaticFS)))
	setupPageRoutes(mux)

	// Serving routes
	slog.Info("App successfully started", "port", ":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return fmt.Errorf("serving api routes: %w", err)
	}

	return nil
}

package http

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reztheperson/ith/internal/web"
)

func Start() error {
	r := mux.NewRouter()

	// Set up routes
	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(web.StaticFS)))

	setupAPIRoutes(r)
	setupPageRoutes(r)

	// Serving routes
	slog.Info("App successfully started", "port", ":8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("serving http routes: %w", err)
	}

	return nil
}

package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reztheperson/ith/internal/web"
)

func Start() error {
	r := mux.NewRouter()

	// Setting up api routes
	//

	// Setting up web routes
	if err := web.Serve(r); err != nil {
		return fmt.Errorf("setting up web routes: %w", err)
	}

	// Serving api routes
	slog.Info("App successfully stared at port :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("serving api routes: %w", err)
	}

	return nil
}

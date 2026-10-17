package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static/* templates/*
var webFS embed.FS

func Serve(r *mux.Router) error {
	// Setup static
	staticFS, err := fs.Sub(webFS, "static")
	if err != nil {
		return err
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// Initialize templates
	if err := InitTemplates(webFS); err != nil {
		return fmt.Errorf("Initializing templates: %w", err)
	}

	r.HandleFunc("/", HomeHandler)

	return nil
}

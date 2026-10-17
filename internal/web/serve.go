package api

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static/* templates/*
var webFS embed.FS

func Serve(r *mux.Router) error {
	if err := InitTemplates(webFS); err != nil {
		return err
	}

	staticFS, err := fs.Sub(webFS, "web")
	if err != nil {
		return err
	}

	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(staticFS)))

	return nil
}

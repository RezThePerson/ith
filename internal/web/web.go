package web

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static/* templates/*
var webFS embed.FS

var Tmpl *template.Template

func InitTemplates() (err error) {
	Tmpl, err = template.ParseFS(webFS, "templates/*.html")
	if err != nil {
		return err
	}

	return nil
}

func SetupStatic(r *mux.Router) error {
	staticFS, err := fs.Sub(webFS, "static")
	if err != nil {
		return err
	}

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	return nil
}

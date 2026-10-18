package web

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
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

func SetupStatic(mux *http.ServeMux) error {
	staticFS, err := fs.Sub(webFS, "static")
	if err != nil {
		return err
	}

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	return nil
}

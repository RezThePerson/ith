package api

import (
	"embed"
	"log/slog"
	"text/template"
)

var tmpl *template.Template

func InitTemplates(webFS embed.FS) (err error) {
	tmpl, err = template.ParseFS(webFS, "templates/*.html")
	if err != nil {
		return err
	}
	slog.Info("templates loaded")
	return nil
}

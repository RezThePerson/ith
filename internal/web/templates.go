package web

import (
	"embed"
	"html/template"
)

var tmpl *template.Template

func InitTemplates(webFS embed.FS) (err error) {
	tmpl, err = template.ParseFS(webFS, "templates/*.html")
	if err != nil {
		return err
	}
	return nil
}

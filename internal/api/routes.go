package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/reztheperson/ith/internal/web/templates"
)

func setupAPIRoutes(mux *http.ServeMux) {

}

func setupPageRoutes(mux *http.ServeMux) {
	mux.Handle("GET /{$}", templ.Handler(templates.Home()))
}

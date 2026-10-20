package http

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/reztheperson/ith/internal/api"
	"github.com/reztheperson/ith/internal/web"
	"github.com/reztheperson/ith/internal/web/templates"
)

func setupAPIRoutes(r *mux.Router) {
	r.HandleFunc("/api/login", api.Login).Methods("GET")
	r.HandleFunc("/api/auth", api.Auth).Methods("GET")
}

func setupPageRoutes(r *mux.Router) {
	r.PathPrefix("/static/").Handler(http.FileServer(http.FS(web.StaticFS)))

	r.Handle("/", templ.Handler(templates.Home())).Methods("GET")
}

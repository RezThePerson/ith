package api

import (
	"log/slog"
	"net/http"

	"github.com/reztheperson/ith/internal/web"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := web.Tmpl.ExecuteTemplate(w, "base", nil)

	if err != nil {
		slog.Error("template execution error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

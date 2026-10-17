package api

import (
	"log/slog"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "base", nil)

	if err != nil {
		slog.Error("template execution error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

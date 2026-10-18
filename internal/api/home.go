package api

import (
	"bytes"
	"log/slog"
	"net/http"

	"github.com/reztheperson/ith/internal/web"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	if err := web.Tmpl.ExecuteTemplate(&buf, "base", nil); err != nil {
		slog.Error("template execution error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := buf.WriteTo(w); err != nil {
		slog.Error("failed to write response", "err", err)
	}
}

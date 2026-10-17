package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() error {
	r := mux.NewRouter()

	
	
	// Serving api routes
	if err := http.ListenAndServe(":8080", r); err != nil {
		return fmt.Errorf("serving api routes: %w", err)
	}

	return nil
}

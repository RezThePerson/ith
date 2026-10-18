package api

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/reztheperson/ith/internal/web"
)

func setupApiRoutes(r *mux.Router) error {

	//

	return nil
}

func setupAppRoutes(r *mux.Router) error {
	// Static files
	if err := web.SetupStatic(r); err != nil {
		return fmt.Errorf("setting up static routes: %w", err)
	}

	// Pages
	r.HandleFunc("/", HomeHandler)

	return nil
}

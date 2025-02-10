package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Build creates a new router
func Build() *mux.Router {
	r := mux.NewRouter()
	return routes.Build(r)
}

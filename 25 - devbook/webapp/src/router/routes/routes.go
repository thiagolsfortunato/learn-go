package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct to define a route
type Route struct {
	URI         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	AuthRequire bool
}

// Build creates a new router
func Build(router *mux.Router) *mux.Router {
	var routes []Route
	routes = append(routes, loginRoutes...)
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		router.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

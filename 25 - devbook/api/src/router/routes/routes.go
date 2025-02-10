package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	AuthRequire bool
}

func Build(r *mux.Router) *mux.Router {
	var routes []Route
	routes = append(routes, userRoutes...)
	routes = append(routes, loginRoutes...)
	routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.AuthRequire {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}

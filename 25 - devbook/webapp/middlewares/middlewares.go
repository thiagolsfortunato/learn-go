package middlewares

import (
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(": %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

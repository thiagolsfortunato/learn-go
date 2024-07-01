package middlewares

import (
	"api/src/auth"
	"api/src/response"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("- %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.CheckToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}

		nextFunc(w, r)
	}
}

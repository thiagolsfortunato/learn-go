package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error returns an API error message
type ApiError struct {
	Error string `json:"error"`
}

// JSON
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ProcessErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var error ApiError
	json.NewDecoder(r.Body).Decode(&error)
	JSON(w, r.StatusCode, error)
}

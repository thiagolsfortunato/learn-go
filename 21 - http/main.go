package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello World"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello from Users"))
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5001", nil))
}
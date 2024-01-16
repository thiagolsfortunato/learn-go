package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main () {
	port := ":5000"

	router := mux.NewRouter()
	
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listening", port)
	log.Fatal(http.ListenAndServe(port, router))
}
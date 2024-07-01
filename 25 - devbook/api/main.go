package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	r := router.Build()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
	log.Println("Listening ", config.Port)
}

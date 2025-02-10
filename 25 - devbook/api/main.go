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

	log.Println("Api listening on", config.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

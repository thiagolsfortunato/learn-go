package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Load()
	utils.LoadTemplates()
	r := router.Build()

	log.Println("WebApp listening on", config.Port)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}

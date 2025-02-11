package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Port int
var ApiUrl string

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("WEBAPP_PORT"))
	if err != nil {
		Port = 8080
	}

	ApiUrl = os.Getenv("API_URL")
	if ApiUrl == "" {
		ApiUrl = "http://localhost:5000"
	}
}

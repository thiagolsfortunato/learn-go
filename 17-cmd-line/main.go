package main

import (
	"fmt"
	"go-cli/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting the go-cli...\n")

	app := app.Build()
	if error := app.Run(os.Args) ; error != nil {
		log.Fatal(error)
	}
	
}
package main

import (
	"fmt"
	// import module utils
	"packages/utils"

	// import external module checkmail
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main.go file")
	utils.Write()

	err := checkmail.ValidateFormat("golearn@gmail.com")
	if err != nil { fmt.Println(err) }
}

package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2) // size of buffer

	channel <- "Hello Wold"
	channel <- "Hello Wold 2"
	channel <- "Hello Wold 3" //dead lock

	message := <- channel
	message2 := <- channel

	fmt.Println(message)
	fmt.Println(message2)
}
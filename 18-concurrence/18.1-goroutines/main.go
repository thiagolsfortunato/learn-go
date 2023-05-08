package main

import (
	"fmt"
	"time"
)

func main() {
	// concurrence != paralellism

	go write("Hello World!") // goroutine
	write("Go app")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// concurrence != paralellism
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func () {
		write("Hello World!") // goroutine
		waitGroup.Done() // -1
	}()

	go func ()  {
		write("Go app")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // wait until waitGroup == 0
}

func write(text string) {
	for i := 0 ; i < 5 ; i++{
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

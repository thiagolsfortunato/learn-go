package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello World")

	for i:=0; i<=10; i++ {
		fmt.Println(<- channel)
	}
}

func write(text string) <- chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received value: %s", text)
			time.Sleep(time.Microsecond * 500)
		}
	}()

	return channel
}
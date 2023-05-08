package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexing(write("Message 1"), write("Message 2"))

	for i:=0; i<=10; i++ {
		fmt.Println(<- channel)
	}
}

func multiplexing(channelIn, channelIn2 <- chan string) <- chan string {
	channelOut := make(chan string)

	go func() {
		for {
			select {
				case msg := <- channelIn:
					channelOut <- msg
				case msg := <- channelIn2:
					channelOut <- msg
			}
		}
	}()

	return channelOut
}

func write(text string) <- chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value: %s", text)
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(200))) // randomic sleeping
		}
	}()

	return channel
}
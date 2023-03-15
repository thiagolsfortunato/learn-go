package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementing I:", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j+= 2 {
		fmt.Println("Incrementing J:", j)
		time.Sleep(time.Millisecond)
	}

	names := [3]string{"Paul", "John", "Steve"}

	for idx, value := range names {
		fmt.Println(idx, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	user := map[string]string {
		"name": "Thiago",
		"lastName": "Fortunato",
	}

	for idx, value := range user {
		fmt.Println(idx, value)
	}

	// Don't works with Structs
}
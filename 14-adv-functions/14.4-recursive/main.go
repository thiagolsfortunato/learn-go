package main

import "fmt"

func fibonacci(index uint) uint {
	if index <= 1 {
		return index
	}

	return fibonacci(index - 2) + fibonacci(index - 1)
}

func main () {

	index := uint(4)
	fmt.Println(fibonacci(index))

	index2 := uint(12)
	for i := uint(0); i < index2 ; i++ {
		fmt.Println(fibonacci(i))
	}
}
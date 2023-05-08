package main

import (
	"fmt"
)

func main() {
	tasks := make(chan int, 50)
	results := make(chan int, 50)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i:=0; i <= 45; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0 ; i <= 45; i++ {
		result := <- results
		fmt.Println(result)
	}
}

func worker (tasks <-chan int, results chan<- int) {
	for num := range tasks {
		results <- fibonacci(num)
	}
}

func fibonacci(index int) int {
	if index <= 1 {
		return index
	}

	return fibonacci(index - 2) + fibonacci(index - 1)
}

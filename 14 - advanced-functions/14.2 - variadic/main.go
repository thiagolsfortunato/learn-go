package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func write(text string, nums...int) {
	for _, num := range nums {
		fmt.Println(text, num)
	}
}

func main() {
	total := sum(1,2,3,4,4,5,5,6,6,24)
	fmt.Println(total)

	write("text", 1,2,3,4,4,5,5,6,56,4,3)
}

// only one variadic parameter per function
// variadic parameter must be the last parameter
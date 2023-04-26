package main

import "fmt"

func invertSignal(num int) int {
	return num * -1
}

func invertSignalPointer(num *int) {
	*num = *num * -1
}

func main() {
	num := 20
	numInverted := invertSignal(num)
	fmt.Println(numInverted)
	fmt.Println(num)

	num2 := 40
	fmt.Println(num2)
	invertSignalPointer(&num2)
	fmt.Println(num2)
}


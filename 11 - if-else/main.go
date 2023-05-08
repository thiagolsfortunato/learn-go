package main

import "fmt"

func main() {
	fmt.Println("Control Structure")

	num := -5

	if num > 15 {
		fmt.Println("Major than 15")
	} else {
		fmt.Println("Minor or equal 15")
	}

	if otherNum := num; otherNum > 0 {
		fmt.Println("Num is greater than zero")
	} else if num < -10 {
		fmt.Println("Num is minor than zero")
	} else {
		fmt.Println("Between 0 and -10")
	}

}
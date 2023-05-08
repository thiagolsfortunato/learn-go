package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")
	}()

	func(text string) {
		fmt.Println(text)
	}("Hello World from Parameter")

	result := func(text string) string{
		return fmt.Sprintf("Received -> %s - %d", text, 10)
	}("Hello World from Parameter")

	fmt.Println(result)
}
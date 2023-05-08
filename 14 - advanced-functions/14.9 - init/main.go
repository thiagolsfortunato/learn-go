package main

import "fmt"

var n int

func main() {
	fmt.Println("main")
	fmt.Println(n)
}

func init() {
	fmt.Println("init")
	n = 10
}

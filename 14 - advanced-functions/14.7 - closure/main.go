package main

import "fmt"

func closure() func(){
	text := "On closure fuction"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "on main function"
	fmt.Println(text)

	funcNew := closure()
	funcNew()
}
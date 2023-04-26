package main

import "fmt"

func one() {
	fmt.Println("executing function one")
}

func two() {
	fmt.Println("executing function two")
}

func isApproved(n1, n2 float32) bool {
	defer fmt.Println(("AVG calculated"))
	fmt.Println("checking if is approved..")

	avg := (n1 + n2) / 2

	if avg >= 6 {
		return true
	}

	return false
}

func main() {
	defer one()
	two()

	fmt.Println(isApproved(6, 3))
}

// defer clauses is executed before the return

package main

import "fmt"

func recoveryFunc() {
	fmt.Println("Trying to recovery the system..")
	if r := recover() ; r != nil {
		fmt.Println("System recovered with success!")
	}
}

func isApproved(n1, n2 float32) bool {
	defer recoveryFunc()
	defer fmt.Println(("AVG calculated"))
	fmt.Println("checking if is approved..")

	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("AVG is equal 6!")
}

func main() {
	fmt.Println(isApproved(6, 6))
	fmt.Println("Post execution")
}
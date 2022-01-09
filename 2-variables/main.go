package main

import "fmt"

func main() {
	// explicit: var <name> <type> = <value>
	var var1 string = "var1"

	// implicit: <name> := <value>
	var2 := "var2"

	fmt.Println(var1, var2)

	// multiple explicit vars
	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	// multiple implict vars
	var5, var6 := "var5", "var6"

	fmt.Println(var3, var4)
	fmt.Println(var5, var6)

	// constant explicit: const <name> <type> = <value>
	const const1 string = "const1"
}

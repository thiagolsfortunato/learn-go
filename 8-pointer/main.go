package main

import "fmt"


func main()  {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	// A pointer holds the memory address of a value.
	var var3 int
	var pointer *int

	fmt.Println(var3, pointer) // initialize Nil

	var3 = 100
	pointer = &var3

	fmt.Println(var3, pointer) // Memory Address
	fmt.Println(var3, *pointer) // This is known as "dereferencing" or "indirecting".

	var3 = 200
	pointer = &var3

	fmt.Println(var3, *pointer) // This is known as "dereferencing" or "indirecting".

}

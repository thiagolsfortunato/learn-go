package main

import "fmt"

func main()  {
	// arithmetics: +, -, *, /, %,
	sum := 1 + 2
	sub := 1 - 2
	mult := 1 * 2
	div := 10 / 2
	restDiv := 10 % 2
	
	fmt.Println(sum, sub, mult, div, restDiv)
	
	// assignment
	var var1 string = "string1"
	var2 := "string2"
	fmt.Println(var1, var2)

	// relational operators: bool
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// logical operators
	varTrue, varFalse := true, false
	fmt.Println(varTrue && varFalse)
	fmt.Println(varTrue || varFalse)
	fmt.Println(!varTrue)
	fmt.Println(!varFalse)

	// unary operators
	num := 10
	num++
	fmt.Println(num)
	num += 10
	fmt.Println(num)
	num--
	fmt.Println(num)
	num *= 2
	fmt.Println(num)

	// ternary operators: not supported :(
	// use if else instead.
	
}
package main

import (
	"fmt"
	"strings"
)

func sum (n1 int8, n2 int8) int8  {
	return n1 + n2
}

func calcMath (n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main()  {
	resultSum := sum(10, 20)
	fmt.Println(resultSum)

	var f = func( txt string) string {
		fmt.Println(txt)
		return strings.ToUpper(txt)
	}

	upper := f("Text func")
	fmt.Println(upper)

	resultCalcSum, resultCalcSub := calcMath(20, 30)
	// resultCalcSum, _ := calcMath(20, 30) Ignore the result of Subtration

	fmt.Println(resultCalcSum, resultCalcSub)
}
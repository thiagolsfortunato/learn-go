package main

import (
	"fmt"
	"reflect"
)


func main()  {
	// Arrays
	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 1
	array1[3] = 1
	fmt.Println(array1)

	array2 := [5]int{0,1,2,3,4}
	fmt.Println(array2)

	array3 := [...]int{0,1,2}
	fmt.Println(array3)

	// Slice
	slice1 := []int{0,1,2,3,4,5,6}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice1))

	var slice2 []int
	slice2 = append(slice2, 1)
	fmt.Println(slice2)

	slice3 := array1[0:2]
	fmt.Println(slice3)

	array1[0] = 10 // Changing value of Array1
	fmt.Println(slice3)

	// Internal Array vs Slices
	slice4 := make([]float32, 10, 15)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacity

	slice5 := make([]float32, 5)
	fmt.Println(len(slice5)) // length
	fmt.Println(cap(slice5)) // capacity

	slice5 = append(slice5, 10)
	fmt.Println(len(slice5)) // length
	fmt.Println(cap(slice5)) // capacity
}

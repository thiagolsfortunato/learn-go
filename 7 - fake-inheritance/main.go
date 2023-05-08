package main

import "fmt"

type person struct {
	name string
	lastName string
	age uint8
	height float32
}

type student struct {
	person
	course string
	college string
}

func main()  {
	p1 := person{"Jhon", "Steve", 20, 1.80}
	fmt.Println(p1)

	s1 := student{p1, "Compute", "Harvard"}
	fmt.Println(s1)
}

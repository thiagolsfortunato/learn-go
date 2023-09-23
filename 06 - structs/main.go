package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	stree string
	number uint
}

func main()  {
	var u user
	u.name = "Fortunato"
	u.age = 30
	fmt.Println(u)

	a := address{"Avenue 1", 123}

	u2 := user{"Jeff", 20, a}
	fmt.Println(u2)

	u3 := user{name: "Jhon"}
	fmt.Println(u3)
}
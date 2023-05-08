package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) getName() {
	fmt.Println(u.name)
}

func (u user) getAge() {
	fmt.Println(u.age)
}

func (u user) isOlder() bool {
	return u.age >= 18
}

func (u *user) haveBirthday() {
	u.age++
}

func main() {
	user1 := user{name: "Thiago", age: 31}
	user1.getName()
	user1.getAge()

	user2 := user{name: "Pedro", age: 18}
	user2.getName()
	fmt.Println(user2.isOlder())

	user1.haveBirthday()
	user1.getAge()
}
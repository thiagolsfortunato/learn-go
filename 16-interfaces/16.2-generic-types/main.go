package main

import (
	"fmt"
)

func generic(interf ...interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String", 1, true)
	generic(1)
	generic(true)

	mess := map[interface{}]interface{} {
		1: "string",
		"String": float32(1.3),
		float32(100): "String",
		true: 32,
	}

	fmt.Println(mess)
}
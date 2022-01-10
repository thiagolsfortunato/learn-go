package main

import (
	"errors"
	"fmt"
)

func main() {

	// Integers: int8, int16, int32, int64
	// default = 0
	var varint8 int8 = 99
	var varint16 int16 = 9999
	var varint32 int32 = 999999999
	var varint64 int64 = 999999999999999999
	var varint0 int

	// "int", assum your chipset architecture
	// int64 because chipset
	var varint int = 999999999999999999

	// unit: unsigned int,
	// also have uint8, uint15, uint32, uint64
	// don't accept negative numbers -1
	var varunit uint = 10000

	// aliases
	// alias for int32: rune
	var varrune rune = 999999999

	// alias for int8: byte
	var varbyte byte = 99

	fmt.Println("int8:", varint8)
	fmt.Println("int16:", varint16)
	fmt.Println("int32:", varint32)
	fmt.Println("int64:", varint64)
	fmt.Println("int0:", varint0)
	fmt.Println("int:", varint)
	fmt.Println("uint:", varunit)
	fmt.Println("rune:", varrune)
	fmt.Println("byte:", varbyte)

	// Float: decimal numbers
	// default = 0
	// explicit declaration: float32 or float64
	var varfloat32 float32 = 1234.56
	var varfloat64 float64 = 1.23456789
	var varfloat0 float64

	// implicit
	// assum your chipset architecture
	// float64 because chipset
	varfloat := 1.23456789

	fmt.Println("float32:", varfloat32)
	fmt.Println("float64:", varfloat64)
	fmt.Println("float0:", varfloat0)
	fmt.Println("float:", varfloat)

	// string
	// default = ""
	var string string = "string"
	fmt.Println("string:", string)

	// convert to int to ASCII number
	// define using single quotes
	char := 'C'
	fmt.Println("char C:", char)

	// boolean: true or false
	// default: false
	var varbool bool = true
	fmt.Println("bool:", varbool)

	// error
	// deafult = <nil>
	var error error = errors.New("Internal Error")
	fmt.Println("error:", error)

}

package main

import (
	"fmt"
	"intro-tests/address"
)

func main() {
	addressType := address.AddressType("Avenue 5th")
	fmt.Println(addressType)
}
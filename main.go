package main

import (
	"fmt"
	"unit-tests/address"
)

func main() {
	var typeAddress string = address.TypeAddress("rua dos testes")
	fmt.Println(typeAddress)
}

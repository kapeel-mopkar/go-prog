package main

import (
	"fmt"
)

func main() {
	var Name, Address = "Kaps", "Goa" // Type inference

	var id, code = 101, 112

	var status = false

	fmt.Printf("%v,%v,%v,%v,%v \n", Name, Address, id, code, status)
}

package main

import (
	"fmt"
)

func main() {
	var Name, Address, id, salary, status = `Kaps`, `Goa`, 1, 999.99, true // Type inference

	fmt.Printf("%v,%v,%v,%v,%v \n", Name, Address, id, salary, status)
}

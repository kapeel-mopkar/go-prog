package main

import (
	"fmt"
)

func main() {
	var Name = "Kaps" // Type inference

	var id = 101

	var status = true

	fmt.Printf("%T-%v,%T-%v,%T-%v \n", Name, Name, id, id, status, status)

}

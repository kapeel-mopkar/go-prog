package main

import (
	"fmt"
)

func main() {
	name := "Kapeel"
	id := 23
	status := true

	fmt.Printf("%T:%v\t%T:%v\t%T:%v\t\n", name, name, id, id, status, status)
}

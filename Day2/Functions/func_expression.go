package main

import (
	"fmt"
)

func main() {
	greet := func() {
		fmt.Println("Hello from greet function")
	}

	fmt.Println(greet) // memory address of function

	greet()
	greet()
	greet()

}

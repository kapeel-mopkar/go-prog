package main

import (
	"fmt"
)

func main() {

	greetings := make([]string, 3, 5)

	// greetings[0] = "hello"
	// greetings[1] = "good morning"
	// greetings[2] = "bonjour"

	// fmt.Println(greetings)

	// greetings[3] = "arigato" // panic: runtime error: index out of range

	greetings = append(greetings, "arigato")
	fmt.Println(greetings)

}

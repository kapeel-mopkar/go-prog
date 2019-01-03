package main

import "fmt"

func main() {
	greeting := []string{
		"java",
		".net",
		"c++",
		"node.js",
		"angularjs",
		"golang",
	}

	fmt.Println(greeting[2:4])

	greeting1 := greeting[2:4]

	fmt.Println(greeting1)

	greeting[2] = "c-plus-plus"

	//greeting and greeting1 refer to same slice/array internally
	fmt.Println(greeting)
	fmt.Println(greeting1)

	fmt.Println(greeting[:2])

	fmt.Println(greeting[3:])

}

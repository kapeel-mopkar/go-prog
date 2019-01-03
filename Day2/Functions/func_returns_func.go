package main

import "fmt"

func main() {

	func1 := func11()
	fmt.Println(func1)
	fmt.Println(func1())
}

func func11() func() int {

	greet := func() int {
		return 1000
	}
	return greet

}

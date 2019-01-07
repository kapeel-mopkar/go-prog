package main

import "fmt"

// var counter int
// func Increment() int {
// 	counter++
// 	return counter
// }

var nextCounter = func() (func() int, func() int) {
	var counter int
	var incrementor = func() int {
		counter++
		return counter
	}

	var decrementor = func() int {
		counter--
		return counter
	}

	return incrementor, decrementor
}

func main() {
	// fmt.Println(Increment())
	// fmt.Println(Increment())
	// fmt.Println(Increment())
	// fmt.Println(Increment())
	incrementor, decrementor := nextCounter()
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	fmt.Println("+++++++++++++++++++++++++")
	incrementor, decrementor = nextCounter()
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	fmt.Println(incrementor())

	fmt.Println(decrementor())
	fmt.Println(decrementor())
	fmt.Println(decrementor())

}

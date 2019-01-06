package main

import "fmt"

func Incrementor() func() int {
	var counter int
	incrementor := func() int {
		counter++
		return counter
	}
	return incrementor
}

func main() {
	counterInc := Incrementor()
	fmt.Println(counterInc())
	fmt.Println(counterInc())
	fmt.Println(counterInc())
	fmt.Println(counterInc())
}

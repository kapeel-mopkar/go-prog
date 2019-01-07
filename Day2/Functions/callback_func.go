package main

import "fmt"

func IterateNumbers(slice1 []int, printNumber func(int)) {
	for _, num := range slice1 {
		printNumber(num)
	}
}

func main() {

	slice1 := []int{1, 2, 3, 4}
	IterateNumbers(slice1, func(num int) {
		fmt.Println(num)
	})
}

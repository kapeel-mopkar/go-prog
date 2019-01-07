package main

import "fmt"

func IterateNumbers(slice1 []int, callback func(num int)) {
	for _, val := range slice1 {
		callback(val)
	}
}
func main() {
	slice1 := []int{1, 2, 3, 4}

	IterateNumbers(slice1, func(num int) {
		fmt.Println(num)
	})
}

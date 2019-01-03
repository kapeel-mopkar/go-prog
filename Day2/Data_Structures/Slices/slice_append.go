package main

import "fmt"

func main() {
	//slice: array without specifying size
	var slice1 = []int{1, 2, 3, 4, 5}

	//appending new element at end of slice
	slice1 = append(slice1, 6, 7, 8, 9)

	// slice refers to array internally
	//slice is of type structure
	//1 pointer to internal array
	//2 length of internal array
	//3 capacity

	fmt.Println(slice1)

}

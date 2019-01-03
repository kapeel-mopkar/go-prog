package main

import (
	"fmt"
)

func main() {
	var slice1 []int
	// slice1[0] = 1 // panic: runtime error: index out of range

	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)

	fmt.Println(len(slice1), cap(slice1), slice1)

	var slice2 = make([]int, 2) // length and capacity would be 2

	slice2[0] = 1
	slice2[1] = 2
	slice2 = append(slice2, 3, 4, 5)
	fmt.Println(len(slice2), cap(slice2), slice2)

}

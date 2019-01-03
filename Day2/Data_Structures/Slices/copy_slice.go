package main

import (
	"fmt"
	"sort"
)

func main() {
	myslice := []int{1, 2, 3, 4}

	myslice1 := make([]int, 4)

	copy(myslice1, myslice)

	fmt.Println(myslice1)

	sort.Ints(myslice)

	//create slice from array
	var array1 = [5]int{1, 2, 3, 4, 5}

	slice1 := array1[:] // all elements in array/slice are included

	fmt.Println(slice1)
}

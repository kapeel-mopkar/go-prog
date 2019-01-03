package main

import "fmt"

func main() {
	//slice: array without specifying size

	var slice1 = []int{1,
		2,
		3,
		4,
		5}

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	for _, v := range slice1 {
		//fmt.Printf("Index:%v\tValue:%v\n", i, v)
		fmt.Printf("Value:%v\n", v)
	}
}

package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Value of i :", i)
	}

	var j = 0
	for j < 10 {
		fmt.Println("Value of j :", j)
		j += 2
	}

	var k = 0
	for {
		k += 1
		if k >= 10 {
			break
		}
		fmt.Println("Value of k :", k)
	}

	var l = 0
	for {
		l++
		if l%2 == 0 {
			continue
		}
		if l >= 20 {
			break
		}
		fmt.Println("Value of l :", l)
	}
}

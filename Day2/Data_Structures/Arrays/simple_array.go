package main

import (
	"fmt"
)

func main() {
	var array1 [3]int // array of type int of len 3 is created

	array1[0] = 10
	array1[1] = 20
	array1[2] = 30

	fmt.Println(array1)

	//declare and initialize

	var array2 = [2]string{"java", "golang"}

	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	var array3 [5]int
	for j := 0; j < len(array3); j++ {

		array3[j] = j + 1
	}
	fmt.Println(array3)

	for index, value := range array3 {
		fmt.Printf("Index:%v\tValue:%v\n", index, value)
	}
}

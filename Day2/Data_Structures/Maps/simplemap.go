package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]int{
		"key1": 10,
		"key2": 20,
		"key3": 30,
	}

	for k, v := range map1 {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	//map2 := map[int]string // Error

	//map2[1] = "Java" // Error

	//make a map

	map2 := make(map[int]string)

	map2[100] = "java"
	map2[200] = ".net"
	map2[300] = "golang"

	fmt.Println(map2)

	//updating map entry
	map2[200] = "Node.js"
	fmt.Println(map2)

	//deleting map entry
	delete(map2, 200)
	fmt.Println(map2)

	//check existance of a key

	if val, existance := map2[300]; existance {
		fmt.Println("Key found", val)
	} else {
		fmt.Println("Key NOT found")
	}

}

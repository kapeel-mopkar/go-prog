package main

import (
	"fmt"
)

func main() {
	var name string = "Jenny"

	//Pattern 1
	// switch name {
	// case "John":
	// 	fmt.Println("John is called")
	// case "Tom":
	// 	fmt.Println("Tom is called")
	// case "Andy":
	// 	fmt.Println("Andy is called")
	// default:
	// 	fmt.Println("Default is executed")
	// }

	//Pattern 2
	// switch name {
	// case "John":
	// 	fmt.Println("John is called")
	// case "Tom":
	// 	fmt.Println("Tom is called")
	// 	fallthrough
	// case "Andy":
	// 	fmt.Println("Andy is called")
	// 	fallthrough
	// case "Ajay":
	// 	fmt.Println("Ajay is called")
	// default:
	// 	fmt.Println("Default is executed")
	// }
	// name = "Sachin"
	// //Pattern 3
	// switch name {
	// case "John", "Jenny":
	// 	fmt.Println("John/Tom is called")
	// case "Sachin", "Ajay":
	// 	fmt.Println("Sachin/Ajay is called")
	// 	//fallthrough
	// default:
	// 	fmt.Println("Default is executed")
	// }

	//Pattern 4 - conditional switch-case
	name = "John"
	switch {
	case len(name) == 4:
		fmt.Println("John/Tom is called")
	case name == "Ken":
		fmt.Println("Ken is called")
		//fallthrough
	case (getSalary() > 1000):
		fmt.Println("Good salary")
	default:
		fmt.Println("Default is executed")
	}

	SwitchOnType("Kaps")
	SwitchOnType(100)
	SwitchOnType(100.50)
}

func getSalary() int {
	return 10000
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown Type")
	}
}

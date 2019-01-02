package main

import (
	"fmt"
)

func main() {
	var i = 10
	if i == 10 && i > 0 {
		fmt.Println("1) if is executed")
	}

	var result = getSalary()

	if result > 999 {
		fmt.Println("2) if is executed")
	} else if result > 2000 {
		fmt.Println("3) else if is executed")
	} else {
		fmt.Println("4) else is executed")
	}

	if result := getSalary(); result > 999 {
		fmt.Println("5) if is executed")
	}
}

func getSalary() int {
	return 10000
}

package main

import (
	"fmt"
)

//global variable declaration
var name string = "Kaps"
var id int = 123
var status bool
var salary float64 = 999.99

//salary = 1999.99 // NOT ALLOWED

func main() {
	//standard variable declaration
	//var name string = "Kaps"
	//var id int = 123
	//var status bool = true
	//var salary float64 = 999.99

	salary = 1999.99
	status = true

	fmt.Println(name, id, status, salary)

}

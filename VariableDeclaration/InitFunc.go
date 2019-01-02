package main

import (
	"Bootcamp/DataAccess"
	"fmt"
	"reflect"
	"strconv"
)

var a, b int

//	Called only once
//	Called before running main func
//		Used for one-time initialization
//		Check for some status, pre-requisite
func init() {
	a = 100
	b = 200
	fmt.Println("1) a and b initialized", a, b)

}

func init() {
	a = 100
	b = 200
	fmt.Println("2) a and b initialized again", a, b)
}

func main() {
	fmt.Println("Main Func called")
	DataAccess.WriteData()

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a))

	//String to Int
	parsedInt, err := strconv.Atoi("1234")
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	fmt.Println("parsedInt", parsedInt)

	//Int to String
	parsedStr := strconv.Itoa(456778)
	fmt.Println("parsedStr", parsedStr)

}

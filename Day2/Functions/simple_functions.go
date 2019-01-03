package main

import (
	"fmt"
	"reflect"
)

func main() {
	// func1("Mukta", "Pune")
	// value1 := func2("")
	// fmt.Println(value1)
	// code, name := func3("")
	// fmt.Println(code, name)

	// amt, status := getBonus("")
	// fmt.Println(amt, status)

	variadicFunc(1, 2, 3, 4, 5, 6, 7, 8)
	variadicFunc(1, 2, 3)
	variadicFunc()

	variadicFunc2("Golang", 1, 1, 2, 3, 4, 5)

	slice1 := []int{10, 20, 30, 40, 50}

	variadicFunc2("Java", 9, slice1...)

	array1 := [2]int{100, 200}

	variadicFunc2("Java", 9, array1[:]...)

	variadicExample(1, "2", true)

}

func func1(name, addr string) {
	fmt.Println(name, addr)
}

func func2(name string) int {
	return 10000
}

func func3(name string) (int, string) {
	code := 10
	lang := "Golang"
	return code, lang
}

func getBonus(name string) (amount int, status bool) {
	amount = 1000
	status = true
	//return amount, status // or below naked return
	return
}

func variadicFunc(values ...int) {
	fmt.Println("variadicFunc1++++++++++++++++++++++++", values, "++++++++++++++++++++++++")
	for i, val := range values {
		fmt.Println(i, val)
	}
}

func variadicFunc2(name string, age int, values ...int) {
	fmt.Println("variadicFunc2++++++++++++++++++++++++", values, "++++++++++++++++++++++++")
	for i, val := range values {
		fmt.Println(name, age, " --- ", i, val)
	}
}

func variadicExample(i ...interface{}) {
	for index, v := range i {
		fmt.Println(index, reflect.TypeOf(v))
	}
}

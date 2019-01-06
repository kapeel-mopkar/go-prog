package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println("Value of i before changeByValue() :", i)
	changeByValue(i)
	fmt.Println("Value of i after changeByValue() :", i)

	fmt.Println("Value of i before changeByReference() :", i)
	changeByReference(&i)
	fmt.Println("Value of i after changeByReference() :", i)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Value of slice1 before changeSlice() :", slice1)
	changeSlice(slice1)
	fmt.Println("Value of slice1 after changeSlice() :", slice1)

	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Value of array1 before changeArray() :", array1)
	changeArray(array1)
	fmt.Println("Value of array1 after changeArray() :", array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Value of array2 before changeArrayRef() :", array2)
	changeArrayRef(&array2)
	fmt.Println("Value of array2 after changeArrayRef() :", array2)

}

func changeByValue(data int) {
	data = 900
}

func changeByReference(data *int) {
	*data = 900
}

func changeSlice(slice1 []int) {
	slice1[0] = 7777
}

func changeArray(array1 [5]int) {
	array1[0] = 7777
}

func changeArrayRef(array1 *[5]int) {
	array1[0] = 7777
}

package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println("Val of I before changeVal() ", i)
	changeVal(i)
	fmt.Println("Val of I after changeVal() ", i)
	fmt.Println("Val of I before changeRef() ", i)
	changeRef(&i)
	fmt.Println("Val of I after changeRef() ", i)

	//how slices are passed by ref
	slice1 := []int{1, 2, 3, 4}
	fmt.Println("Val of slice1 before changeSlice() ", slice1)
	changeSlice(slice1)
	fmt.Println("Val of slice1 after changeSlice() ", slice1)

	//how arrays are passed by value
	array1 := [4]int{1, 2, 3, 4}
	fmt.Println("Val of array1 before changeArray() ", array1)
	changeArray(array1)
	fmt.Println("Val of array1 after changeArray() ", array1)

	//how arrays are passed by ref
	array2 := [4]int{1, 2, 3, 4}
	fmt.Println("Val of array2 before changeArrayRef() ", array2)
	changeArrayRef(&array2)
	fmt.Println("Val of array2 after changeArrayRef() ", array2)

}

func changeVal(data int) {
	data = 900
}

func changeRef(data *int) {
	*data = 900 // dereferencing pointer
}

func changeSlice(slice1 []int) {
	slice1[0] = 987654
}

func changeArray(array1 [4]int) {
	array1[0] = 987654
}

func changeArrayRef(array2 *[4]int) {
	array2[0] = 987654
}

package main

import "fmt"

func main() {
	var slice2 = make([]int, 2, 8) // length(len) and capacity(cap) would be 2

	fmt.Println("Length: ", len(slice2), ", Capacity: ", cap(slice2))
	slice2[0] = 100
	slice2[1] = 200

	fmt.Println("Length: ", len(slice2), ", Capacity: ", cap(slice2))

	slice2 = append(slice2, 300)

	fmt.Println("Length: ", len(slice2), ", Capacity: ", cap(slice2))

	slice2 = append(slice2, 400)
	fmt.Println("Length: ", len(slice2), ", Capacity: ", cap(slice2))

	slice2 = append(slice2, 500, 600)
	fmt.Println("Length: ", len(slice2), ", Capacity: ", cap(slice2))

	slice3 := append(slice2, 700)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)

	slice2[4] = 7777
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)

}

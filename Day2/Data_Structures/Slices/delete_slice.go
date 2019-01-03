package main

import "fmt"

func main() {
	myslice := []string{"Monday", "Tuesday"}

	otherslice := []string{"Wednesday", "Thursday", "Friday"}

	myslice = append(myslice, otherslice...)

	fmt.Println(myslice)

	//delete element at second index
	//slicing + app

	myslice = append(myslice[:2], myslice[3:]...)
	fmt.Println(myslice)
}

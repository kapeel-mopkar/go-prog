package main

import "fmt"

type Product struct {
	Name string
	ID   int
}

func main() {
	var prd1 = Product{Name: "TV", ID: 100}
	fmt.Println("Before modifyStruct(): ", prd1)
	modifyStruct(prd1)
	fmt.Println("After modifyStruct(): ", prd1)

	fmt.Println("Before modifyStructByRef(): ", prd1)
	modifyStructByRef(&prd1)
	fmt.Println("After modifyStructByRef(): ", prd1)

}

func modifyStruct(prd Product) {
	prd.Name = "Samsung TV"
}

func modifyStructByRef(prd *Product) {
	prd.Name = "Samsung TV"
}

package main

import "fmt"

type Product struct {
	name string
	id   int
}

func main() {
	var product1 = Product{name: "Cellphone", id: 101}
	fmt.Println("Before modifyStructure() :", product1.name, product1.id)
	modifyStructure(product1)
	fmt.Println("After modifyStructure() :", product1.name, product1.id)

	fmt.Println("Before modifyStructureRef() :", product1.name, product1.id)
	modifyStructureRef(&product1)
	fmt.Println("After modifyStructureRef() :", product1.name, product1.id)

}

func modifyStructure(p Product) {
	fmt.Println("\tProduct being updated ...", p.id)
	p.name = "Cellphone v1.1"
}

func modifyStructureRef(p *Product) {
	fmt.Println("\tProduct being updated ...", p.id)
	p.name = "Cellphone v1.1"
}

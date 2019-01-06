package main

import "fmt"

type Item struct {
	Name string
	ID   int
}

func main() {

	//Value type
	item1 := Item{"Laptop", 9876}
	fmt.Println("Value Type: ", item1)
	//Creating Ref Type 1
	item2 := new(Item)
	item2.ID = 987
	item2.Name = "Mobile"
	fmt.Println("Reference Type 1: ", item2)

	//Creating Ref Type 2
	item3 := &Item{"Mouse", 22}
	fmt.Println("Reference Type 2 Before modifyItem(): ", item3)
	modifyItem(item3)
	fmt.Println("Reference Type 2 After modifyItem(): ", item3)

}

func modifyItem(item *Item) {
	item.Name += " (Samsung)"
}

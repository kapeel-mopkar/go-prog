package main

import "fmt"

type Item struct {
	Name string
	Id   int
}

func main() {
	item1 := Item{"Laptop", 1111}
	fmt.Println("Value Type : ", item1)
	item2 := new(Item)
	item2.Name = "Printer"
	item2.Id = 2222
	fmt.Println("Ref Type 1 : ", item2)

	item3 := &Item{"Mouse", 3333}
	fmt.Println("Ref Type 2 [Before calling modifyItemStruct()]: ", item3)
	modifyItemStruct(item3)
	fmt.Println("Ref Type 2 [After calling modifyItemStruct()]: ", item3)

}

func modifyItemStruct(item *Item) {
	item.Name += " - SAMSUNG"
}

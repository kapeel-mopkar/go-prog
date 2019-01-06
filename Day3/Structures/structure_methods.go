package main

import "fmt"

type Person struct {
	Name string
	ID   int
}

//Method for Person struct
//func with receiver method
func (p Person) GetPersonDetails(caller string) {
	fmt.Println("Data from method", p.ID, p.Name, "Caller-", caller)
}

//GetPersonDetails - Normal function
//Allowed in file along with method with same name
func GetPersonDetails() {
	fmt.Println("GetPersonDetails")
}

func main() {
	p := Person{"Kapeel Mopkar", 23}

	p.GetPersonDetails("Main")
	GetPersonDetails()
}

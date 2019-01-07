package main

import "fmt"

type Person struct {
	Name string
	ID   int
}

//Method for Person struct
//func with receiver method
//Receiver are of 2 types, - Value type and Ref type
//Below receiver is Value Type receiver
func (p Person) GetPersonDetails1(caller string) {
	fmt.Println("Data from method", p.ID, p.Name, "Caller-", caller)
}

//Below receiver is Reference Type receiver
func (p *Person) GetPersonDetails2(caller string) {
	fmt.Println("Data from method", p.ID, p.Name, "Caller-", caller)
	p.ID = 999
}

func main() {
	p1 := Person{"Kapeel Mopkar", 23}

	p1.GetPersonDetails1("Main")

	fmt.Println("Person 1, Before GetPersonDetails2() :", p1)
	p1.GetPersonDetails2("Main")
	fmt.Println("Person 1, After GetPersonDetails2() :", p1)

	p2 := new(Person)
	p2.Name = "Sourav"
	p2.ID = 100

	fmt.Println("Person 2, Before GetPersonDetails2() :", p2)
	p2.GetPersonDetails2("Main")
	fmt.Println("Person 2, After GetPersonDetails2() :", p2)

}

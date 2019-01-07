package main

import "fmt"

//Structure:
//		can hold members of different datatypes
//		Employee mgmt: EmpName, EmpNum, Code, Salary etc

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{
		firstName: "Kapeel",
		lastName:  "Mopkar",
		id:        1234,
	}
	fmt.Println(emp1.id, ":", emp1.firstName, ":", emp1.lastName)

	var emp2 Employee
	fmt.Println(emp2.id, ":", emp2.firstName, ":", emp2.lastName)

	var emp3 Employee
	emp3.firstName = "Sam"
	emp3.lastName = "Pitts"
	emp3.id = 2345

	fmt.Println(emp3.id, ":", emp3.firstName, ":", emp3.lastName)

	emp4 := Employee{
		"Kaps",
		"Mopkar",
		1235,
	}
	fmt.Println(emp4.id, ":", emp4.firstName, ":", emp4.lastName)
}

package main

import "fmt"

type Employee struct {
	fname string
	lname string
	id    int
}

func main() {

	//Pattern 1
	emp1 := Employee{fname: "Kapeel", lname: "Mopkar", id: 16828}
	fmt.Println(emp1.id, " - '", emp1.fname, "', '", emp1.lname, "'")

	//Pattern 2
	var emp2 Employee
	fmt.Println(emp2.id, " - '", emp2.fname, "', '", emp2.lname, "'")

	//Pattern 3
	emp3 := Employee{}
	emp3.fname = "Sam"
	emp3.lname = "Mopkar"
	emp3.id = 12345
	fmt.Println(emp3.id, " - '", emp3.fname, "', '", emp3.lname, "'")

	//Pattern 4
	emp4 := Employee{"Sachin", "Tendulkar", 10}
	fmt.Println(emp4.id, " - '", emp4.fname, "', '", emp4.lname, "'")

}

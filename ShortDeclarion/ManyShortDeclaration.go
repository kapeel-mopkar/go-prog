package main

import (
	"fmt"

	_ "github.com/golang/example/stringutil" // Blank identifier
)

func main() {
	name, id, status := "Kapeel", 23, true

	fmt.Printf("%T:%v\t%T:%v\t%T:%v\t\n", name, name, id, id, status, status)

	var name1 = "PSL"

	//Only for standard variable declaration
	var (
		a, b int
		c, d string
	)

	_ = name1               // Blank identifier
	_, _, _, _ = a, b, c, d // Blank identifier

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//stringutil.Reverse("TEST")
}

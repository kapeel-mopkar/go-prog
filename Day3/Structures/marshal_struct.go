package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First                 string
	Last                  string
	Age                   int
	notExportedSecretCode int
}

func main() {
	p1 := person{"James", "Bond", 35, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println("Byte String of Marshalled Person - ", bs)
	fmt.Printf("Print Type : %T \n", bs)
	fmt.Println(string(bs))

}

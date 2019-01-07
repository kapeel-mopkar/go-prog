package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is a self executing function")
	}()
}

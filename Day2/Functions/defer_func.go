package main

import "fmt"

func main() {
	defer foo()
	bar()
	fmt.Println("main is called")

	fmt.Println("Main", fetchData())
}

func foo() {
	fmt.Println("foo is called")
}

func bar() {
	fmt.Println("bar is called")
}

func fetchData() string {
	defer closeConnection()
	status := true
	if status {
		fmt.Println("Connected")
		return "Connected"
	} else {
		fmt.Println("Connection Failed")
		return "Connection Failed"
	}
}

func closeConnection() {
	fmt.Println("Connections are closed")
}

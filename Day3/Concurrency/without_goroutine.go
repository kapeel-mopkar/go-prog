package main

import (
	"fmt"
	"time"
)

func main() {

	func1()
	func2()
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func1", i)
	}

}

func func2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func2", i)
	}
}

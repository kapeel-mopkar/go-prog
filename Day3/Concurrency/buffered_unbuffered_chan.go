package main

import "fmt"

func main() {
	ch1 := make(chan int, 2) // buffer is 2

	ch1 <- 12 // error if unbuffered
	ch1 <- 13
	//ch1 <- 14 // error becoz buffer is 2, this is 3rd value added on chan
	fmt.Println("Value retrieved from channel", <-ch1)
	fmt.Println("Value retrieved from channel", <-ch1)
	ch1 <- 14
	ch1 <- 15
	fmt.Println("Value retrieved from channel", <-ch1)
	fmt.Println("Value retrieved from channel", <-ch1)
}

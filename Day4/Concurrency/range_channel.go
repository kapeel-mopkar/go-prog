package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go ProcessData(ch1)

	for value := range ch1 {
		fmt.Println("Value from channel :", value)
	}

}

func ProcessData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // informs receivers that channel values are no longer being sent on channel
}

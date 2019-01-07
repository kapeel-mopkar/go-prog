package main

import (
	"fmt"
	"time"
)

func Server1(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 1000
	close(ch)
}

func Server2(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 2000
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Server1(ch1)
	go Server2(ch2)

	mergeChannels(ch1, ch2)
}

func mergeChannels(chs ...chan int) {
	for value, data := range chs {
		fmt.Println("Values are", value, <-data)
	}
}

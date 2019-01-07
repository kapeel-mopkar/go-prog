package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int) // unbuffered channel

	wg.Add(2)
	go func1(ch1) // go routine
	go func2(ch1) // go routine
	wg.Wait()
}

// func1 going to put something on channel
func func1(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("func1 putting value into channel", i)
		ch1 <- i
	}
	wg.Done()
}

// func2 will retrieve something from channel
func func2(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("func2 retrieving value from channel", <-ch1)
	}
	wg.Done()
}

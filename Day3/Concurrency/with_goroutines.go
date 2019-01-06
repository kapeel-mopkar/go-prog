package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func1() // go routine
	go func2() // go routine
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func1", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000)
		fmt.Println("i from func2", i)
	}
	wg.Done()
}

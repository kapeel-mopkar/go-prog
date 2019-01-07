package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutx sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo :")
	go incrementor("Bar :")
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {

		mutx.Lock() // Avoiding race conditions using mutex.Lock/Unlock
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		mutx.Unlock() // Avoiding race conditions using mutex.Lock/Unlock
		fmt.Println(s, i, "counter: ", counter)
	}
	wg.Done()

}

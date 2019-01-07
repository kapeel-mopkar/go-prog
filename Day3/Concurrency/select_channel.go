package main

import (
	"fmt"
	"time"
)

func Server1(ch chan string) {
	time.Sleep(30 * time.Second)
	ch <- "Response from Server1"
	//fmt.Println("Response from Server1")
}

func Server2(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Response from Server2"
	//fmt.Println("Response from Server2")
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go Server1(ch1)
	go Server2(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
			return
		case msg2 := <-ch2:
			fmt.Println(msg2)
			return
		default:
			time.Sleep(4 * time.Second)
			fmt.Println("Waiting for response from Server")
		}
	}

}

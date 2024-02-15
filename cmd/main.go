package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go sender1(ch1)
	go sender2(ch2)
	select {
	case value := <-ch1:
		fmt.Printf("Got from channel one: %v \n", value)
	case value := <-ch2:
		fmt.Printf("Got from channel two: %v \n", value)
	case <-time.After(1 * time.Second):
		fmt.Printf("Time out \n")
	}

	time.Sleep(1 * time.Second)
}

func sender1(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 1
}

func sender2(ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "Qingzhi"
}

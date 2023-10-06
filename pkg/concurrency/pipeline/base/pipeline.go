package base

import (
	"fmt"
)

func Gen(nums ...int) <-chan int {
	fmt.Println("Start gen func ...")
	defer fmt.Println("End gen func ...")
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func Sq(in <-chan int) <-chan int {
	fmt.Println("Start sq func ...")
	defer fmt.Println("End sq func ...")
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

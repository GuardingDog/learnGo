package cancel

import (
	"context"
	"fmt"
	"sync"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
				fmt.Printf("stage one send : %v \n", n)
			case <-ctx.Done():
				fmt.Printf("OH CANCEL STAGE ONE \n")
				return
			}
		}
	}()
	return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
				fmt.Printf("stage two send : %v \n", n*n)
			case <-ctx.Done():
				fmt.Printf("OH CANCEL STAGE TWO \n")
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-ctx.Done():
				fmt.Printf("OH CANCEL STAGE THREE \n")
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

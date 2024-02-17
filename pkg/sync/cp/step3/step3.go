package step3

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	Count int
}

func (c *Counter) Inc() {
	time.Sleep(10 * time.Millisecond)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Count++
}

func (c *Counter) Value() int {
	return c.Count
}

func assertCount(c *Counter, want int) {
	if c.Value() == want {
		fmt.Println("Assert SUCCESS!!!")
	} else {
		fmt.Printf("Assert FAILED, want: %v, got: %v \n", want, c.Value())
	}
}

func Assert() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	assertCount(c, 10)
}

func Assert1000() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	assertCount(c, 1000)
}

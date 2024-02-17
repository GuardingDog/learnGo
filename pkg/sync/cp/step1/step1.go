package step1

import (
	"fmt"
	"time"
)

type Counter struct {
	Count int
}

func (c *Counter) Inc() {
	time.Sleep(10 * time.Millisecond)
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
	for i := 0; i < 10; i++ {
		c.Inc()
	}

	assertCount(c, 10)
}

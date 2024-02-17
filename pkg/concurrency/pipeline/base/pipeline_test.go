package base

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	ch1 := Gen(1, 2, 3)
	// fan-out: Distribute the work across two goroutines
	out1 := Sq(ch1)
	out2 := Sq(ch1)
	// fan-in: Consume fan-in result
	for o := range merge(out1, out2) {
		fmt.Printf("Got result: %v \n", o)
	}
}

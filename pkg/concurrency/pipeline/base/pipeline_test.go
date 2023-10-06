package base

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	c := Gen(1, 2, 3, 4, 5, 6, 7, 8, 9)
	out := Sq(c)
	for n := range out {
		fmt.Println(n)
	}
}

func TestPipeline2(t *testing.T) {
	for n := range Sq(Sq(Sq(Gen(2, 3)))) {
		fmt.Println(n)
	}
}

package fan

import (
	"fmt"
	"testing"

	"learngo/pkg/concurrency/pipeline/base"
)

func TestPipeline3(t *testing.T) {
	in := base.Gen(2, 3)

	c1 := base.Sq(in)
	c2 := base.Sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

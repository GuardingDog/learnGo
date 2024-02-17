package cancel

import (
	"context"
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	in := gen(ctx, 2, 3, 4, 5)
	c1 := sq(ctx, in)
	c2 := sq(ctx, in)
	out := merge(ctx, c1, c2)
	fmt.Println(<-out)
}

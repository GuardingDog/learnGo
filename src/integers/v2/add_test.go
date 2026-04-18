package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}

package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

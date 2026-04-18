package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 8})
	want := []int{3, 8}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	got := SumAllTails([]int{1, 2}, []int{0, 9, 1})
	want := []int{2, 10}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

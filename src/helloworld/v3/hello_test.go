package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("QingZhi")
	want := "Hello, World! QingZhi"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
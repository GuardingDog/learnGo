package main

import "testing"

func TestAssertBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Assert()
	}
}

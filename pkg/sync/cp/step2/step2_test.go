package step2

import "testing"

func TestAssert(t *testing.T) {
	Assert1000()
}

func BenchmarkAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Assert()
	}
}

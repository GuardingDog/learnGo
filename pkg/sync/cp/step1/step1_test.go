package step1

import "testing"

func TestAssert(t *testing.T) {
	Assert()
}

func BenchmarkAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Assert()
	}
}

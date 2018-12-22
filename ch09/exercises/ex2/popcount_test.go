package popcount

import (
	"testing"
)

func BenchmarkPopcount(b *testing.B) {
	x := uint64(141)
	for i := 0; i < b.N; i++ {
		Popcount(x)
	}
}

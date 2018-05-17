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

func BenchmarkPopcountFor(b *testing.B) {
	x := uint64(141)
	for i := 0; i < b.N; i++ {
		PopcountFor(x)
	}
}

func BenchmarkPopcountShift(b *testing.B) {
	x := uint64(141)
	for i := 0; i < b.N; i++ {
		PopcountShift(x)
	}
}

func BenchmarkPopcountClear(b *testing.B) {
	x := uint64(141)
	for i := 0; i < b.N; i++ {
		PopcountClear(x)
	}
}

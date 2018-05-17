package ex3

import (
	"testing"
)

func BenchmarkEfficient(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		Efficient(s)
	}
}

func BenchmarkInefficient(b *testing.B) {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		Inefficient(s)
	}
}

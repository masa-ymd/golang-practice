package main

import (
	"testing"
)

func BenchPipeline(b *testing.B) {
	in, out := pipeline(10000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

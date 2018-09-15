package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/masa-ymd/golang-practice/ch02/exercises/ex3"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%d\n", popcount.Popcount(binary.BigEndian.Uint64(c1[:])^binary.BigEndian.Uint64(c2[:])))
}

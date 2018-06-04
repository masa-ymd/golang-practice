package main

import (
	"bytes"
	"fmt"
	"sort"
)

type buf []byte

func (b buf) Len() int {
	return len(b)
}

func (b buf) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b buf) Less(i, j int) bool {
	return b[i] < b[j]
}

func main() {
	fmt.Println(compare("hoge", "ghoe"))
	fmt.Println(compare("hoge", "ugaf"))
}

func compare(s1, s2 string) bool {
	b1 := buf(bytes.NewBufferString(s1).Bytes())
	b2 := buf(bytes.NewBufferString(s2).Bytes())
	sort.Sort(b1)
	sort.Sort(b2)
	return bytes.Equal(b1, b2)
}

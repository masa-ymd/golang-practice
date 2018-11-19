package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	return strconv.Itoa(t.value)
}

func main() {
	t := &tree{1, nil, nil}
	fmt.Printf("%s\n", t)
}

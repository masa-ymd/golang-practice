package main

import (
	"fmt"
)

func main() {
	fmt.Println(join(",", "a", "b", "c"))
}

func join(sep string, vals ...string) string {
	var s string
	for i, v := range vals {
		if i == 0 {
			s = v
		} else {
			s = s + sep + v
		}
	}

	return s
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(expand("ab$cd", toUpper))
}

func expand(s string, f func(string) string) string {
	i := strings.Index(s, "$")

	if i != -1 {
		return f(s[i:])
	}

	return ""
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

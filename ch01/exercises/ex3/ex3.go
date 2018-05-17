package ex3

import (
	"fmt"
	"strings"
)

func Inefficient(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Efficient(args []string) {
	fmt.Println(strings.Join(args, " "))
}

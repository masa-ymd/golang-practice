package main

import (
	"bufio"
	"fmt"
	"os"
	//"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		s := in.Text()
		counts[s]++
	}
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileCounts := make(map[string]map[string]int)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
		countLines(os.Stdin, fileCounts, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// 第一引数で出力先を指定可能
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileCounts, counts)
			f.Close()
		}
	}

	for filename, counts := range fileCounts {
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, filename)
			}
		}
	}
}

func countLines(f *os.File, fileCounts map[string]map[string]int, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileCounts[f.Name()] = counts
	}
}

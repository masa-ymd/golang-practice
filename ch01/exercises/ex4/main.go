package main

import (
	"bufio"
	"fmt"
	"os"
)

type cnt struct {
	num       int
	filenames []string
}

func main() {
	counts := make(map[string]*cnt)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, c := range counts {
		if c.num > 1 {
			fmt.Printf("%s\t%d\t%s\n", c.filenames, c.num, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*cnt) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if val, ok := counts[input.Text()]; ok {
			c := val
			c.num++
			if !arrayContains(c.filenames, f.Name()) {
				c.filenames = append(c.filenames, f.Name())
			}
		} else {
			c := cnt{
				num:       1,
				filenames: []string{f.Name()},
			}
			counts[input.Text()] = &c
		}
	}
}

func arrayContains(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}

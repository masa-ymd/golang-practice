package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func main() {
	var c *WordCounter
	fmt.Println(c.Write([]byte("hoge fuga aaa")))
	var l *LineCounter
	fmt.Println(l.Write([]byte("aaa\nbbb")))
}

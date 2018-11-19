package main

import (
	"bytes"
	"fmt"
	"io"
)

type byteCounter struct {
	w       io.Writer
	written int64
}

func (c *byteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.written += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &byteCounter{w, 0}
	return c, &c.written
}

func main() {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	fmt.Println(fmt.Fprintf(c, "hoge hoge"))
	fmt.Println(*n)
}

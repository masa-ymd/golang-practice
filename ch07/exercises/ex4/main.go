package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type stringReader struct {
	s string
}

func (s *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, s.s)
	s.s = s.s[n:]
	if len(s.s) == 0 {
		err = io.EOF
	}
	return n, err
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}

func main() {
	s := "<html><body><p>hi</p></body></html>"
	n, _ := html.Parse(NewReader(s))
	fmt.Println(n.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.Data)
}

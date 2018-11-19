package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r        io.Reader
	n, limit int
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p[:l.limit])
	l.n += n
	if l.n >= l.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}

func main() {
	r := LimitReader(strings.NewReader("hoge hoge"), 4)
	b := &bytes.Buffer{}
	s, _ := b.ReadFrom(r)
	fmt.Println(s)
}

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(commafloat("-10000000.1234"))
}

func commafloat(s string) string {
	var buf bytes.Buffer
	l := strings.Split(s, ".")
	b := []byte(l[0])
	if string(b[0]) == "-" || string(b[0]) == "+" {
		buf.WriteByte(b[0])
		b = b[1:]
	}
	buf.WriteString(comma(string(b)))
	if len(l) == 2 {
		buf.WriteRune('.')
		buf.WriteString(l[1])
	}
	return buf.String()
}

func comma(s string) string {
	var buf bytes.Buffer
	b := []byte(s)
	var cnt int
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(b[i])
		cnt++
		if cnt == 3 {
			buf.WriteByte(',')
			cnt = 0
		}
	}
	b2 := buf.Bytes()
	var buf2 bytes.Buffer
	for i := len(b2) - 1; i >= 0; i-- {
		buf2.WriteByte(b2[i])
	}
	return buf2.String()
}

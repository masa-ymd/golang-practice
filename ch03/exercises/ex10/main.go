package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("10000000"))
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

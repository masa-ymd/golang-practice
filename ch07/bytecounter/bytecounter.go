package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // int(len) を　ByteCouter形に変換
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // カウンターリセット
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name) // c.Writeが呼び出され、hello以降を書き込む
	fmt.Println(c)
}

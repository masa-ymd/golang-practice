package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFileはバイトスライスを返す
		// 入力全体を一気に読み込む
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// バイトスライスなのでcast
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	// bufio.Scanner, ioutil.ReadFile, ioutil.WriteFileは引数で渡された
	// *os.FileのRead, Writeメソッドを使用しているが、この低レベルルーチンを
	// プログラマが使用する必要はほぼない
	// bufio, ioutilの高レベル関数を使うほうが良い
}

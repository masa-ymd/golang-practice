package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	// 初期化の方法はいろいろあるが、以下の２形式が推奨
	// s := ""
	// var s string
	// range 文字列、スライスの値の範囲を繰り返す
	// インデックスとインデックスに対応する値を返す
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = " "
	for i := 1; i < len(os.Args); i++ {
        // os.Args[] は引数を格納するスライス
        // インデックス0はコマンド自身の名前
        // goのスライスのインデックス範囲指定は最初を含み最後を含まない
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}

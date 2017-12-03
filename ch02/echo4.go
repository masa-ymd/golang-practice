package main

import (
	"flag"
	"fmt"
	"strings"
)

// 引数はフラグ名、デフォルト値、メッセージ
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
        // 改行コードを表示
		fmt.Println()
	}
}

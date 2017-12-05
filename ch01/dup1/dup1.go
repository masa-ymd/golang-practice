package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
    // Scanner型(入力を読み込み行や単語に分解)を生成
    // ストリームモードで、必要に応じて入力が読み込まれる
	input := bufio.NewScanner(os.Stdin)
	// Scanは呼び出す毎に次の行を返す　なくなるまでtrue
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			// fでおわるものは%d, %sなどのフォーマット変換を使える
			fmt.Printf("%d\t%s\n", n, line)
			// lnでおわるものは引数が%v(任意の値に対する自然なフォーマット)
			// に変換される
		}
	}
}

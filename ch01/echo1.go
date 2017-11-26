package main

import (
	"fmt"
	"os"
)

func main() {
    // 変数宣言時に、初期化の値が明示的に指定されない時は
    // それぞれの方に対するゼロ値が設定される
    // intは0,stringは""など nilではない
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

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type IntSet struct {
}

func (*IntSet) String() string { return "hoge" }

func main() {
	var w io.Writer // write メソッドが必要
	// writeメソッドを持つ値はインターフェースに代入可能
	w = os.Stdout
	w = new(bytes.Buffer)

	var rwc io.ReadWriteCloser

	w = rwc // インターフェースにインターフェースを代入することも可能

	fmt.Printf("%T\n", w)

	// var _ = IntSet{}.String() 値がないためコンパイルエラー

	var s IntSet
	var _ = s.String() // 暗黙的にポインタに変換するため、メソッド呼び出し可能

	var _ fmt.Stringer = &s // StringメソッドがあるのでStringerインターフェースに代入可能
	// var _ fmt.Stringer = s // IntSetにはStringメソッドがないのでエラー

	// インターフェースが満足する関係性をドキュメント化して主張する方法
	var _ io.Writer = (*bytes.Buffer)(nil) // nilを*byte.Buffer型に変換
}

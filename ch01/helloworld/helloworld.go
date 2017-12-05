package main

import (
	"fmt"
)

// import後のプログラムの構成は、だいたい
// 関数(func), 変数(var), 定数(const), 型(type)
// になる

// 関数宣言は func 関数名(パラメータのリスト(mainは空)) 結果のリスト(mainは空) {}
func main() {
	var x int
	// +の後なら改行OK
	x = x +
		x
	fmt.Println("Hello, 世界.")
}

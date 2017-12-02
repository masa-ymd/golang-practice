package main

import "fmt"

// パッケージレベルの宣言
const boilingF = 212.0

func main() {
	// 関数レベルの宣言
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	// インデックスに対応するASCII文字コードが表示される
	fmt.Printf("%d %d %[1]b %[2]b\n", s[0], s[6])
	fmt.Println(s[0:2])
	s2 := "ハローワールド"
	fmt.Printf("%d %d %d %[1]b %[2]b %[3]b\n", s2[0], s2[1], s2[2])
	fmt.Println(s2[0:3])
	// 生文字リテラル->エスケープシーケンスは処理されない
	fmt.Println(`aaa\naaa`)
	s3 := "\x51\x42\xe2\x9a\xa0"
	fmt.Println(s3)
	fmt.Printf("%d %b\n", len(s3), s3[0])
	s4 := "\101\102"
	fmt.Println(s4)
	fmt.Printf("%d %b\n", len(s4), s4[0])
	// 世界の世はutf-8 16進数エンコードでe4b896
	fmt.Println("\xe4\xb8\x96")
	// unicodeで4e16
	fmt.Println("\u4e16")
	fmt.Println("\U00004e16")

	s5 := "Hello, 世界"
	for i := 0; i < len(s5); {
        // runeをutf8に変換
		r, size := utf8.DecodeRuneInString(s5[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

    s6 := "プログラム"
    // utf8エンコード
    fmt.Printf("% x\n", s6)
    r := []rune(s6)
    // rune(unicode)
    fmt.Printf("%x\n", r)

    fmt.Println(string(65))

}

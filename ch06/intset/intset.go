package main

import (
	"bytes"
	"fmt"
)

// 1スライスに0～63まで格納　次のスライスは64から始まる
// 値がはいっていたらbitが立つ
type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Printf("Has %b, %b, %b, %b, %b\n", word, bit, s.words[word], 1<<bit, s.words[word]&(1<<bit))
	// &&の右の項目がpanicにならないように、左の項目でチェック
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		fmt.Printf("add %b %b %d\n", word, bit, len(s.words))
		s.words = append(s.words, 0)
	}
	fmt.Printf("%b\n", s.words[word])
	// すでに立っているbitが消えないようにor算
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					// 数値の区切りはスペース
					buf.WriteByte(' ')
				}
				// i=スライス単位なので*64 jはbit bitが立っている時だけ、bufに数値を書き込み
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x IntSet
	x.Add(1)
	fmt.Println(x.Has(1))
	fmt.Println(x.Has(2))
	x.Add(0)
	fmt.Println(x.Has(0))
	x.Add(64)
	x.Add(65)
	//	fmt.Println(x.String())

	//	y.Add(9)
	//	y.Add(42)
	//	fmt.Println(y.String())

	//	x.UnionWith(&y)
	fmt.Println(x.String())

	//	fmt.Println(x.Has(9), x.Has(123))

	//	fmt.Println(&x)
	//	fmt.Println(x.String())
	//	fmt.Println(x)
}

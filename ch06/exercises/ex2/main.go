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
	//fmt.Printf("Has %b, %b, %b, %b, %b\n", word, bit, s.words[word], 1<<bit, s.words[word]&(1<<bit))
	// &&の右の項目がpanicにならないように、左の項目でチェック
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		//fmt.Printf("add %b %b %d\n", word, bit, len(s.words))
		s.words = append(s.words, 0)
	}
	// すでに立っているbitが消えないようにor算
	s.words[word] |= 1 << bit
	//fmt.Printf("%08b\n", s.words[word])
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
	for i, ward := range s.words {
		if ward == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if ward&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var cnt int
	for _, w := range s.words {
		for w != 0 {
			cnt++
			//fmt.Printf("%08b %v, %08b %v\n", w, w, w-1, w-1)
			w &= w - 1
			//fmt.Printf("%08b %v\n", w, w-1)
		}
	}

	return cnt
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	for i, w := range s.words {
		s.words[i] &^= w
	}
}

func (s *IntSet) Copy() *IntSet {
	res := &IntSet{}
	res.words = make([]uint64, len(s.words))
	for i, w := range s.words {
		res.words[i] = w
	}
	return res
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		word, bit := x/64, uint(x%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		// すでに立っているbitが消えないようにor算
		s.words[word] |= 1 << bit
	}
}

func main() {
	var x IntSet
	x.AddAll(1, 2, 3, 4)
	fmt.Println(x.String())
}

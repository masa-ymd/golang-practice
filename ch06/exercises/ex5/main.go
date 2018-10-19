package main

import (
	"bytes"
	"fmt"
)

const size = 32 << (^uint(0) >> 63)

// 1スライスに0～63まで格納　次のスライスは64から始まる
// 値がはいっていたらbitが立つ
type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/size, uint(x%size)
	//fmt.Printf("Has %b, %b, %b, %b, %b\n", word, bit, s.words[word], 1<<bit, s.words[word]&(1<<bit))
	// &&の右の項目がpanicにならないように、左の項目でチェック
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/size, uint(x%size)
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
		for j := 0; j < size; j++ {
			if ward&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", size*i+j)
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
	word, bit := x/size, uint(x%size)
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	for i, w := range s.words {
		s.words[i] &^= w
	}
}

func (s *IntSet) Copy() *IntSet {
	res := &IntSet{}
	res.words = make([]uint, len(s.words))
	for i, w := range s.words {
		res.words[i] = w
	}
	return res
}

func main() {
	var x IntSet
	var y IntSet
	x.Add(0)
	x.Add(2)
	x.Add(4)
	y.Add(0)
	fmt.Println(x.Len())
	x.Remove(2)
	fmt.Println(x.String())
	x.Clear()
	fmt.Println(x.String())
	z := y.Copy()
	fmt.Println(z.String())
	//fmt.Println(x.Has(1))
	//fmt.Println(x.Has(2))
	y.UnionWith(&x)
	//fmt.Println(y)

}

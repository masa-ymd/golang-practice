package main

import (
	"fmt"
	"sort"
)

// string は一度初期化すると変更不可のため、byte型を利用
type Sentence []byte

func (snt Sentence) Len() int {
	return len(snt)
}

func (snt Sentence) Less(i, j int) bool {
	return snt[i] < snt[j]
}

func (snt Sentence) Swap(i, j int) {
	snt[i], snt[j] = snt[j], snt[i]
}

func IsPalindrome(s sort.Interface) bool {
	loopLimit := s.Len()/2 + 1
	length := s.Len() - 1
	for i := 0; i < loopLimit; i++ {
		if !s.Less(i, length-i) && !s.Less(length-i, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	var s, t Sentence
	s = []byte("abcddcba")
	t = []byte("abcdxcba")
	fmt.Println(IsPalindrome(s), IsPalindrome(t))
}

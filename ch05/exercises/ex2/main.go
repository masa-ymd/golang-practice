package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	tok := html.NewTokenizer(os.Stdin)

	for elm, cnt := range elementCount(tok) {
		fmt.Printf("elm=%s cnt=%d\n", elm, cnt)
	}
}

func elementCount(t *html.Tokenizer) map[string]int {
	cnt := make(map[string]int)
	for {
		type_ := t.Next()
		if type_ == html.ErrorToken {
			break
		}
		elm, _ := t.TagName()
		cnt[string(elm)]++
	}
	return cnt
}

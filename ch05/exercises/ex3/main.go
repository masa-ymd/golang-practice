package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	tok := html.NewTokenizer(os.Stdin)

	for _, t := range textNode(tok) {
		fmt.Printf("%s\n", t)
	}
}

func textNode(t *html.Tokenizer) []string {
	var txt []string
	var stflg int
Tokenize:
	for {
		switch t.Next() {
		case html.StartTagToken:
			n, _ := t.TagName()
			if string(n) == "script" || string(n) == "style" {
				stflg = 1
			}
		case html.TextToken:
			if stflg == 1 {
				continue
			}
			txt = append(txt, string(t.Text()))
		case html.EndTagToken:
			stflg = 0
		case html.ErrorToken:
			break Tokenize
		}
	}
	return txt
}

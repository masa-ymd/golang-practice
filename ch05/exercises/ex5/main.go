package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	switch n.Type {
	case html.TextNode:
		in := bufio.NewScanner(strings.NewReader(n.Data))
		in.Split(bufio.ScanWords)
		for in.Scan() {
			words++
		}
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childWords, childImages := countWordsAndImages(c)
		words += childWords
		images += childImages
	}

	return
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Errorf("%v\n", err)
		}
		fmt.Printf("words=%d, images=%d", words, images)
	}
}

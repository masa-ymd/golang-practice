package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
	// 独自の型を定義し、recoverから返される値がこの型かどうか判断することでpanicから回復すべきか判断
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		// パニックなし
		case bailout{}:
			// 予期されたパニック
			err = fmt.Errorf("multiple title elements")
		default:
			// 予期しないパニック
			panic(p)
		}
	}()

	// 2つ以上の空ではないtitleを見つけたら再帰から抜け出させる
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // recoverの返り値にbailout{}が格納される
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	_, _ = soleTitle(doc)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		title(url)
	}
}

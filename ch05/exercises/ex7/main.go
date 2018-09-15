package main

import (
	"fmt"
	"os"
	//"strings"

	"golang.org/x/net/html"

	"github.com/masa-ymd/golang-practice/ch05/exercises/ex7/forEachNode"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	// 引数と返り値の型を指定すれば関数を引数として渡せる
	forEachNode.ForEachNode(doc, forEachNode.StartElement, forEachNode.EndElement)
}

/*
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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		end := ">"
		if n.FirstChild == nil {
			end = "/>"
		}
		var attrs []string
		for _, a := range n.Attr {
			attrs = append(attrs, fmt.Sprintf("%s=%s", a.Key, a.Val))
		}
		attrStr := strings.Join(attrs, " ")
		fmt.Printf("%*s<%s %s%s\n", depth*2, "", n.Data, attrStr, end)
		depth++
	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) != 0 {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
	// テストを食わせていったん全部n.Dataを表示してみる
	//html.DoctypeNode
	//html.ErrorNode
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
*/

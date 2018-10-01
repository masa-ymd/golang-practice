package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `<html><body><h1>hoge</h1><img></img></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, d := range images {
		fmt.Printf("%v\n", d.Data)
	}
	for _, d := range headings {
		fmt.Printf("%v\n", d.Data)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	var search func(doc *html.Node)
	search = func(doc *html.Node) {
		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				for _, n := range name {
					if c.Data == n {
						nodes = append(nodes, c)
					}
				}
			}
			search(c)
		}
	}
	search(doc)
	return nodes
}

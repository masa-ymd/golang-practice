package forEachNode

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var depth int
var short int

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func StartElement(n *html.Node) {
	if n.Type == html.ElementNode {
		end := ">"
		if n.FirstChild == nil {
			short = 1
			end = "/>"
		}
		var attrs []string
		for _, a := range n.Attr {
			attrs = append(attrs, fmt.Sprintf("%s=%s", a.Key, a.Val))
		}
		attrStr := strings.Join(attrs, " ")
		if attrStr != "" {
			attrStr = " " + attrStr
		}
		fmt.Printf("%*s<%s%s%s\n", depth*2, "", n.Data, attrStr, end)
		depth++
	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) != 0 {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if short == 1 {
			short = 0
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

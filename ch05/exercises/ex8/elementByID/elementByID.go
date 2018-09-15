package elementByID

import (
	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if pre(n) == false {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := forEachNode(c, pre, post)
		if res != nil {
			return res
		}
	}

	if post != nil {
		post(n)
	}

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {

		if n.Type != html.ElementNode {
			return true
		}

		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return false
			}
		}

		return true
	}

	return forEachNode(doc, pre, nil)
}

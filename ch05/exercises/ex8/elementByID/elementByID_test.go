package elementByID

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(t *testing.T) {
	s := `<html><head></head><body><p id="example">Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	id := "example"
	res := ElementByID(doc, id)

	var match bool

	if res != nil {
		for _, a := range res.Attr {
			if a.Key == "id" && a.Val == id {
				match = true
				break
			}
		}
	}

	if !match {
		t.Errorf("%s\n", "id value doesn't match")
	}
}

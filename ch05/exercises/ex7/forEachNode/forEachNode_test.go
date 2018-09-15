package forEachNode

import (
	"strings"

	"golang.org/x/net/html"
)

func ExampleForEachNode() {
	s := `<html><head></head><!-- test --><body><p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul></body></html>`
	doc, _ := html.Parse(strings.NewReader(s))
	ForEachNode(doc, StartElement, EndElement)

	// Output:
	// <html>
	//   <head/>
	//   <!-- test -->
	//   <body>
	//     <p>
	//       Links:
	//     </p>
	//     <ul>
	//       <li>
	//         <a href=foo>
	//           Foo
	//         </a>
	//       </li>
	//       <li>
	//         <a href=/bar/baz>
	//           BarBaz
	//         </a>
	//       </li>
	//     </ul>
	//   </body>
	// </html>
}

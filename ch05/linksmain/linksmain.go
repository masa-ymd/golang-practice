package main

import (
	"fmt"
	"os"

	"github.com/masa-ymd/golang-practice/ch05/links"
)

func main() {
	for _, url := range os.Args[1:] {
		links, _ := links.Extract(url)
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

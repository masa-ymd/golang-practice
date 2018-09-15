package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/masa-ymd/golang-practice/ch05/links"
)

func breathFirst(f func(item string) []string, worklist []string) {
	var worklist2 []string
	seen := make(map[string]bool)
	for _, d := range worklist {
		worklist2 = append(worklist2, d)
		for len(worklist2) > 0 {
			items := worklist2
			worklist2 = nil
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					worklist2 = append(worklist2, f(item)...)
				}
			}
		}
		cnt++
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	urlArray := strings.Split(url, "/")
	if urlArray[2] == domains[cnt] {
		dir := strings.Join(urlArray[2:], "/")
		if _, err := os.Stat(dir); err != nil {
			os.MkdirAll(dir, 0777)
		}
		res, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		defer res.Body.Close()
		file, err := os.Create(dir + "/" + "index.html")
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()
		io.Copy(file, res.Body)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var domains []string
var cnt int

func main() {
	for _, d := range os.Args[1:] {
		ds := strings.Split(d, "/")
		domains = append(domains, ds[2])
	}
	fmt.Println(domains)
	breathFirst(crawl, os.Args[1:])
}

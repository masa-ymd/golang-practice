package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth > *maxDepth {
		return
	}
	list, err := visit(url)
	if err != nil {
		log.Print(err)
	}
	for _, link := range list {
		lock.Lock()
		if seen[link] {
			lock.Unlock()
			continue
		}
		seen[link] = true
		lock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
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

func linkNodes(n *html.Node) []*html.Node {
	var links []*html.Node
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, n)
		}
	}
	forEachNode(n, visitNode, nil)
	return links
}

func linkURLs(linkNodes []*html.Node, base *url.URL) []string {
	var urls []string
	for _, n := range linkNodes {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			// ignore bad and non-local URLs
			if err != nil {
				log.Printf("skipping %q: %s", a.Val, err)
				continue
			}
			if link.Host != base.Host {
				//log.Printf("skipping %q: non-local host", a.Val)
				continue
			}
			urls = append(urls, link.String())
		}
	}
	return urls
}

func rewriteLocalLinks(linkNodes []*html.Node, base *url.URL) {
	for _, n := range linkNodes {
		for i, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			if err != nil || link.Host != base.Host {
				continue // ignore bad and non-local URLs
			}
			// Clear fields so the url is formatted as /PATH?QUERY#FRAGMENT
			link.Scheme = ""
			link.Host = ""
			link.User = nil
			a.Val = link.String()
			n.Attr[i] = a
		}
	}
}

func visit(rawurl string) (urls []string, err error) {
	fmt.Println(rawurl)
	resp, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GET %s: %s", rawurl, resp.Status)
	}

	u, err := base.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if base.Host != u.Host {
		log.Printf("not saving %s: non-local", rawurl)
		return nil, nil
	}

	var body io.Reader
	contentType := resp.Header["Content-Type"]
	if strings.Contains(strings.Join(contentType, ","), "text/html") {
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", u, err)
		}
		nodes := linkNodes(doc)
		urls = linkURLs(nodes, u) // Extract links before they're rewritten.
		rewriteLocalLinks(nodes, u)
		b := &bytes.Buffer{}
		err = html.Render(b, doc)
		if err != nil {
			log.Printf("render %s: %s", u, err)
		}
		body = b
	}
	err = save(resp, body)
	return urls, err
}

func save(resp *http.Response, body io.Reader) error {
	u := resp.Request.URL
	filename := filepath.Join(u.Host, u.Path)
	if filepath.Ext(u.Path) == "" {
		filename = filepath.Join(u.Host, u.Path, "index.html")
	}
	err := os.MkdirAll(filepath.Dir(filename), 0777)
	if err != nil {
		return err
	}
	fmt.Println("filename:", filename)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	if body != nil {
		_, err = io.Copy(file, body)
	} else {
		_, err = io.Copy(file, resp.Body)
	}
	if err != nil {
		log.Print("save: ", err)
	}
	// Check for delayed write errors, as mentioned at the end of section 5.8.
	err = file.Close()
	if err != nil {
		log.Print("save: ", err)
	}
	return nil
}

var (
	maxDepth = flag.Int("depth", 1, "depth of crawl")
)

var seen = make(map[string]bool)
var lock sync.Mutex
var base *url.URL

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	// 未探索のリンクを取得するため20個のクローラーのゴルーチンを起動
	for _, link := range flag.Args() {
		u, err := url.Parse(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid url: %s\n", err)
		}
		base = u
		wg.Add(1)
		// バッファなしチャネルは誰かが読み取らないと先に進めないので別のゴルーチンで実行
		go crawl(link, 0, &wg)
	}
	wg.Wait()
}

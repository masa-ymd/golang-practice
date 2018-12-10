package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/masa-ymd/golang-practice/ch05/links"
)

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth > *maxDepth {
		return
	}
	fmt.Println(url, depth)
	list, err := links.Extract(url)
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

var (
	maxDepth = flag.Int("depth", 1, "depth of crawl")
)

var seen = make(map[string]bool)
var lock sync.Mutex

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	fmt.Println(*maxDepth)

	// 未探索のリンクを取得するため20個のクローラーのゴルーチンを起動
	for _, link := range flag.Args() {
		wg.Add(1)
		// バッファなしチャネルは誰かが読み取らないと先に進めないので別のゴルーチンで実行
		go crawl(link, 0, &wg)
	}
	wg.Wait()
}

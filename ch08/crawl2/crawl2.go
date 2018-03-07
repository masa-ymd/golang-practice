package main

import (
	"fmt"
	"log"
	"os"

	"github.com/masa-ymd/golang-practice/ch05/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // トークンを獲得
	list, err := links.Extract(url)
	<-tokens // トークンを開放
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // worklistへの送信待ちの数

	n++
	// バッファなしチャネルは誰かが読み取らないと先に進めないので別のゴルーチンで実行
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/masa-ymd/golang-practice/ch05/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // 重複を含むURLリスト
	unseenLinks := make(chan string) // 重複を含まないURLリスト

	// バッファなしチャネルは誰かが読み取らないと先に進めないので別のゴルーチンで実行
	go func() { worklist <- os.Args[1:] }()

	// 未探索のリンクを取得するため20個のクローラーのゴルーチンを起動
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				// バッファなしチャネルは誰かが読み取らないと先に進めないので別のゴルーチンで実行
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

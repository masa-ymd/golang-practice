package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // go ルーチンを開始
	}
    f, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
	for range os.Args[1:] { // 引数分だけgoルーチンが生成されている
		// fmt.Println(<-ch) // チャネルから受信
        f.Write(([]byte)(<-ch))
	}
	f.Write(([]byte)(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // チャネルへエラー内容を送信
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

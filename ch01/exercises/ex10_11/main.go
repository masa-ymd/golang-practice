package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, _ := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	for i, url := range os.Args[1:] {
		go fetch(url, ch, file, i)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elaplsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, out io.Writer, n int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	time.Sleep(time.Duration(n) * time.Second)

	nbytes, err := io.Copy(out, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

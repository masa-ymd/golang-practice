package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	p := make(chan int)
	var cnt int64
	start := time.Now()

	go func() {
		p <- 1
		for {
			cnt++
			// <-pでチャネルから値を取り出す
			p <- <-p
		}
	}()
	go func() {
		for {
			p <- <-p
		}
	}()

	c := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 1バイト読み込み
		close(c)
	}()
	<-c

	fmt.Println(float64(cnt)/float64(time.Since(start))*1e9, "round trips per second")

}

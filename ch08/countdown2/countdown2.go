package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 1バイトの入力待ち
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	select {
	case <-time.After(10 * time.Second):
	// 何もしない
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}

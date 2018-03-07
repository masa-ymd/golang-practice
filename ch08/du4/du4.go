package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	// キャンセル後は何もせずに終わる
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20) // 同時20までdirentsの平行稼働を許す

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // トークンの取得
	case <-done:
		return nil // キャンセルされた
	}
	defer func() { <-sema }() // トークンの開放
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait() // すべての操作が終わるまで待つ
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 1バイト読み込み
		close(done)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop: // ラベル breakで指定するとforを抜ける
	for {
		select {
		case <-done:
			// 既存のゴルーチンが受信待ちにならないようfilesizeを読み取り空にする
			for range fileSizes {
				// 何もしない
			}
		case size, ok := <-fileSizes: // チャネルがクローズされるとforloopを抜ける
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

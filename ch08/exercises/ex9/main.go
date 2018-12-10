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

type SizeResponse struct {
	root int
	size int64
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dir string, n *sync.WaitGroup, root int, sizeResponse chan<- SizeResponse) {
	defer n.Done()
	// キャンセル後は何もせずに終わる
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, root, sizeResponse)
		} else {
			sizeResponse <- SizeResponse{root, entry.Size()}
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
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
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

	sizeResponse := make(chan SizeResponse)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, i, sizeResponse)
	}
	go func() {
		n.Wait() // すべての操作が終わるまで待つ
		close(sizeResponse)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // 1バイト読み込み
		close(done)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop: // ラベル breakで指定するとforを抜ける
	for {
		select {
		case <-done:
			// 既存のゴルーチンが受信待ちにならないようfilesizeを読み取り空にする
			for range sizeResponse {
				// 何もしない
			}
		case sr, ok := <-sizeResponse: // チャネルがクローズされるとforloopを抜ける
			if !ok {
				break loop
			}
			nfiles[sr.root]++
			nbytes[sr.root] += sr.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

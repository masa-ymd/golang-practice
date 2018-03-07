package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/masa-ymd/golang-practice/ch08/thumbnail/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Print(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
			fmt.Println(thumb)
		}(f)
	}

	// closer
	// wgが0になるのをまってチャネルをクローズ
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Println(total)

	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

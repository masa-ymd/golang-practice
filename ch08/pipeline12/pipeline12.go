package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // チャネルが空の場合はループから抜ける
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}

package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		// closeすることで、受信側にこれ以上がデータを送信されないことを伝える
		close(naturals)
	}()

	go func() {
		// チャネルでもrangeが利用できる
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

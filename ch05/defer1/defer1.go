package main

import (
	"fmt"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // x==0ならパニック
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	f(3)
}

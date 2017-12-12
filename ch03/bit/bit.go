package main

import (
	"fmt"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	// %bで2進数表記
	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)

	fmt.Printf("and: %08b\n", x&y)
	fmt.Printf("or: %08b\n", x|y)
	fmt.Printf("xor: %08b\n", x^y)
	fmt.Printf("and not: %08b\n", x&^y)

	fmt.Printf("x<<1: %08b\n", x<<1)
	fmt.Printf("x>>1: %08b\n", x>>1)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // and計算でbitが立つのは1, 5
			fmt.Printf("%08b\n", x&(1<<i))
		}
	}

    // 8進数
    o := 0666
    // %oは8進数
    fmt.Printf("%d %[1]o %#[1]o\n", o)
    // 16進数
    X := int64(0xdeadbeef)
    fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", X)

    ascii := 'a'
    unicode := 'あ'
    fmt.Printf("%d %[1]c %[1]q\n", ascii)
    fmt.Printf("%d %[1]c %[1]q\n", unicode)
}

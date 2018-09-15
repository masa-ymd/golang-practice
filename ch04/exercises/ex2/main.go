package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	h = flag.String("h", "sha256", "hash function")
)

func main() {
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	txt := input.Text()

	switch *h {
	case "sha256":
		fmt.Println(sha256.Sum256([]byte(txt)))
	case "sha512":
		fmt.Println(sha512.Sum512([]byte(txt)))
	case "sha384":
		fmt.Println(sha512.Sum384([]byte(txt)))
	default:
		fmt.Errorf("unknown parameter")
	}
}

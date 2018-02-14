package main

import (
	"flag"
	"fmt"

	"github.com/masa-ymd/golang-practice/ch07/tempconv2"
)

var temp = tempconv2.CelsiusFlag("temp", 20.0, "the temparature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

package main

import (
	"fmt"

	"github.com/masa-ymd/golang-practice/ch06/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}

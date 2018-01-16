package main

import (
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	// ...intパラメータは関数内ではスライスだが、引数の際は型が異なる
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

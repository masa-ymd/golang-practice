package main

import (
	"fmt"

	"github.com/masa-ymd/golang-practice/ch06/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	distanceFromP := p.Distance // メソッド値
	fmt.Println(distanceFromP(q))

	distance := geometry.Point.Distance // メソッド式
	fmt.Println(distance(p, q))         // レシーバに該当するものを第一引数として渡して実行させる
}

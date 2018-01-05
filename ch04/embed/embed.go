package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

// 無名フィールドの埋め込み
type Circle struct {
	Point
	Radius int
}

// 無名フィールドの埋め込み
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8 // x.Circle.Point.X = 8 と同じ
	fmt.Printf("%#v\n", w)
	// 宣言時は無名構造体の省略は不可。以下の2つのやり方がある
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
	// w.circle.point.Xのように途中が公開されていないものも、w.Xであればアクセス可能
}

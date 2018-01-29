package main

import (
	"fmt"
	"image/color"

	"github.com/masa-ymd/golang-practice/ch06/geometry"
)

/*
type Point struct {
	X, Y float64
}
*/

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	fmt.Println(cp.Point.X)
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{geometry.Point{1, 1}, red}
	q := ColoredPoint{geometry.Point{5, 4}, blue}

	// PointのDistanceメソッドはColoredPointへ格上げされている
	fmt.Println(p.Distance(q.Point))

}

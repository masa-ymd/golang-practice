package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // キャンバスの大きさ
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲 -xyrange～+xyrange
	xyscale       = width / 2 / xyrange // x,y単位当たりの画素数
	zscale        = height * 0.4        // z単位当たりの画素数
	angle         = math.Pi / 6         // x, y軸の角度（=30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin30°, cos30°

type zFunc func(x, y float64) float64

func svg(f zFunc) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j, f)
			bx, by, berr := corner(i, j, f)
			cx, cy, cerr := corner(i, j+1, f)
			dx, dy, derr := corner(i+1, j+1, f)
			if aerr != nil || berr != nil || cerr != nil || derr != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f zFunc) (float64, float64, error) {
	// マス目 (i, j) の角の点 (x, y) を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := f(x, y)

	// (x, y, z)を2-D　SVGキャンパス(sx, sy)へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(z, 0) {
		return 0, 0, errors.New("z is Inf")
	}
	return sx, sy, nil
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

var (
	t = flag.String("t", "", "picture type")
)

func main() {
	flag.Parse()
	var f zFunc
	switch *t {
	case "egg":
		f = eggbox
	case "saddle":
		f = saddle
	default:
		os.Exit(2)
	}
	svg(f)
}

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	// width, height = 600, 320            // キャンバスの大きさ
	cells   = 100  // 格子のマス目の数
	xyrange = 30.0 // 軸の範囲 -xyrange～+xyrange
	// xyscale = width / 2 / xyrange // x,y単位当たりの画素数
	// zscale  = height * 0.4        // z単位当たりの画素数
	angle = math.Pi / 6 // x, y軸の角度（=30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin30°, cos30°

func svg(width, height int, clr string) string {
	var res string
	res = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j, width, height)
			bx, by, berr := corner(i, j, width, height)
			cx, cy, cerr := corner(i, j+1, width, height)
			dx, dy, derr := corner(i+1, j+1, width, height)
			if aerr != nil || berr != nil || cerr != nil || derr != nil {
				continue
			}
			res += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
				"style='stroke: %s;  fill: white; stroke-width: 0.7'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, clr)
		}
	}
	res += fmt.Sprint("</svg>")

	return res
}

func corner(i, j, width, height int) (float64, float64, error) {
	xyscale := float64(width) / 2 / float64(xyrange)
	zscale := float64(height) * 0.4

	// マス目 (i, j) の角の点 (x, y) を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := f(x, y)

	// (x, y, z)を2-D　SVGキャンパス(sx, sy)へ等角的に投影
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	if math.IsInf(z, 0) {
		return 0, 0, errors.New("z is Inf")
	}
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)  // (0, 0)からの距離
	return math.Sin(r) / r // 半径を1にするためrで割るとsinθが高さになる
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width := 600
	height := 320
	clr := "#000000"

	if w := r.FormValue("width"); w != "" {
		width, _ = strconv.Atoi(w)
	}
	if h := r.FormValue("height"); h != "" {
		height, _ = strconv.Atoi(h)
	}
	if c := r.FormValue("clr"); c != "" {
		clr = "#" + c
	}

	fmt.Fprint(w, svg(width, height, clr))
}

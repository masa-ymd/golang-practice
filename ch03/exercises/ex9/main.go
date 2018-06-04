package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	width := 1024
	if wdh := r.FormValue("width"); wdh != "" {
		width, _ = strconv.Atoi(wdh)
	}
	height := 1024
	if hgt := r.FormValue("height"); hgt != "" {
		height, _ = strconv.Atoi(hgt)
	}
	createMandelbrot(w, width, height)
}

func createMandelbrot(w io.Writer, width, height int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点(px, py)は複素数値zを表している
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
	//fmt.Fprintf(w, "%s", "hoge")
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}

		}
	}
	return color.Black
}

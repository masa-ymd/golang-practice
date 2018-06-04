package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var colorPattern = []color.RGBA{
	{170, 57, 57, 255},
	{170, 108, 57, 255},
	{34, 102, 102, 255},
	{45, 136, 45, 255},
}

var chosenColors = map[complex128]color.RGBA{}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)
	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			subPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					// complexは実部と虚部から複素数を生成
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, newton(z))
				}
			}
			img.Set(px, py, avg(subPixels))
		}
	}
	png.Encode(os.Stdout, img)
}

// x_n+1 = x_n - f(x)/f'(x)
// f'(x) = 4 * z^3
// f(x) = z^4 - 1
func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
		z = z - (z*z*z*z-1)/(z*z*z*4)
		// 解の誤差が小さくなったら終了
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			// realとimagは複素数の値から、それぞれ実数と虚数部分を取り出す
			root := complex(round(real(z)), round(imag(z)))
			c, ok := chosenColors[root]
			if !ok {
				chosenColors[root] = colorPattern[0]
				c = colorPattern[0]
				colorPattern = colorPattern[1:]
			}
			return color.RGBA{c.R, c.G, c.B, c.A}
		}
	}
	return color.Black
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		tr, tg, tb, ta := c.RGBA()
		r += uint16(tr / uint32(n))
		g += uint16(tg / uint32(n))
		b += uint16(tb / uint32(n))
		a += uint16(ta / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func round(f float64) float64 {
	return math.Floor(f + 0.5)
}

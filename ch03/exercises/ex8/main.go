package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点(px, py)は複素数値zを表している
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
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

func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}

		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	r := big.NewFloat(0)
	i := big.NewFloat(0)
	for n := uint8(0); n < iterations; n++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		t1 := new(big.Float).Mul(r, r)
		t2 := new(big.Float).Mul(i, i)
		t3 := new(big.Float).Sub(t1, t2)
		z1 := big.NewFloat(real(z))
		r = new(big.Float).Add(t3, z1)
		two := big.NewFloat(2)
		t4 := new(big.Float).Mul(two, r)
		t5 := new(big.Float).Mul(t4, i)
		z2 := big.NewFloat(imag(z))
		i = new(big.Float).Add(t5, z2)
		rt, _ := r.Float64()
		it, _ := i.Float64()
		if cmplx.Abs(complex(rt, it)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	r := big.NewRat(0, 1)
	i := big.NewRat(0, 1)
	for n := uint8(0); n < iterations; n++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		t1 := new(big.Rat).Mul(r, r)
		t2 := new(big.Rat).Mul(i, i)
		t3 := new(big.Rat).Sub(t1, t2)
		z1 := new(big.Rat).SetFloat64(real(z))
		r = new(big.Rat).Add(t3, z1)
		two := big.NewRat(2, 1)
		t4 := new(big.Rat).Mul(two, r)
		t5 := new(big.Rat).Mul(t4, i)
		z2 := new(big.Rat).SetFloat64(imag(z))
		i = new(big.Rat).Add(t5, z2)
		rt, _ := r.Float64()
		it, _ := i.Float64()
		if cmplx.Abs(complex(rt, it)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

package main

import (
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	format := flag.String("f", "jpg", "output format")
	flag.Parse()
	img, _, _ := image.Decode(os.Stdin)
	switch *format {
	case "jpg":
		jpeg.Encode(os.Stdout, img, nil)
	case "png":
		png.Encode(os.Stdout, img)
	case "gif":
		gif.Encode(os.Stdout, img, nil)
	}
}

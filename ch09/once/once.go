package main

import (
	"image"
	"sync"
)

var loadIconOnce sync.Once // Onceはミューテックスと初期化が行われたかどうかを記録するboolanで構成される
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func Icon(name string) image.Image {
	loadIconOnce.Do(loadIcons) // Do は引数で初期化関数を受け取り、その関数について排他制御を行う
	return icons[name]
}

func loadIcon(name string) image.Image {
	var a image.Image
	return a
}

func main() {
	Icon("spades.png")
}

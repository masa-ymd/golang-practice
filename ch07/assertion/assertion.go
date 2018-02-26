package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	// 具象型で判定を行った場合、その型にcastした値が得られる
	f, ok := w.(*os.File)
	fmt.Printf("%T, %v, %v\n", f, ok, f == os.Stdout)

	// インタフェースで判定を行った場合は、メソッドを保持しているか検査し、そのインターフェース型を得られる
	// 以下の例では　read methodも使えるようになる
	rw, ok2 := w.(io.ReadWriter)
	fmt.Printf("%T, %v %v\n", rw, ok2, rw == os.Stdout)

	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Println(os.IsNotExist(err))

}

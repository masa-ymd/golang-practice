package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// ファイルを閉じるが、Copyのエラーがあれば優先する
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	defer resp.Body.Close()
	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// サーバはそれぞれのリクエストに対して、
    // 別のgoルーチンでhandlerを実行する
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handlerに動作を書く
// *http.Requestにリクエスト内容が保存されている
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

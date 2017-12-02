package main

import (
	"fmt"
	"log"
	"net/http"
    "sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handlerに動作を書く
// *http.Requestにリクエスト内容が保存されている
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// countを返す
func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
    mu.Unlock()
}

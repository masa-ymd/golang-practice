package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// Dialはネットワークnet上のリモードアドレスaddrへ接続
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

// Must が付いているものは error を返す代わりに panic を発生させる。
func mustCopy(dst io.Writer, src io.Reader) {
	// サーバからの出力を標準出力にコピー→ストリーミング出力
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

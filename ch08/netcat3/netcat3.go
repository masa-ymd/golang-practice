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
	// バックグラウンドのゴルーチンの状況を管理するためのチャネルを生成
	// 同期目的の場合チャンネルの要素型をstruct{}にすることで用途を強調する
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	// サーバに標準入力を渡す
	mustCopy(conn, os.Stdin)
	conn.Close()
	// チャネルからデータの受信を待つ（受信した内容は捨てる）
	<-done
}

// Must が付いているものは error を返す代わりに panic を発生させる。
func mustCopy(dst io.Writer, src io.Reader) {
	// サーバからの出力を標準出力にコピー→ストリーミング出力
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

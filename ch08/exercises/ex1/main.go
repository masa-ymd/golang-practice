package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var (
		p = flag.String("port", "8080", "access port")
	)
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*p)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Acceptは接続要求を待ち、接続を表すnet.Connオブジェクトを返す
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // 個々の呼び出しが、個別のgorutine内で実行される
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // クライアントとの接続が切れたときなど
		}
		time.Sleep(1 * time.Second)
	}
}

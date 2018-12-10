package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

// chan<- 送信専用チャネル
// <-chan 受信専用チャネル
func scan(c io.Reader, txt chan<- string) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		txt <- input.Text()
	}
	if input.Err() != nil {
		log.Printf("%s\n", input.Err())
	}
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	txt := make(chan string)
	timer := time.NewTimer(10 * time.Second)
	go scan(c, txt)
loop:
	for {
		select {
		case <-timer.C:
			timer.Stop()
			break loop
		case stxt := <-txt:
			timer.Reset(10 * time.Second)
			wg.Add(1)
			// 複数入力を受け付ける
			go echo(c, stxt, 1*time.Second, &wg)
		}
	}
	wg.Wait()
	c.Close()
}

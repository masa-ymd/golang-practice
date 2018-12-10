package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name    string
	message chan<- string // 送信メッセージ用チャネル
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // 接続中のクライアントを格納するmap
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.message <- msg
			}
		case cli := <-entering: // 基底型がチャネル型のclientが渡される
			clients[cli] = true
			for c := range clients {
				cli.message <- c.name
			}
		case cli := <-leaving: // 基底型がチャネル型のclientが渡される
			delete(clients, cli)
			close(cli.message)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{who, ch}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

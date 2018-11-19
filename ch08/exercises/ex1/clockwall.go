package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name string
	host string
}

func (c *clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "End")
	if s.Err() != nil {
		fmt.Errorf("%s", s.Err())
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		a := strings.Split(arg, "=")
		c := &clock{a[0], a[1]}
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			fmt.Errorf("%s\n", err)
		}
		defer conn.Close()
		go c.watch(os.Stdout, conn)
	}
	for {
		time.Sleep(time.Minute)
	}
}

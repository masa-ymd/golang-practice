package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		if err := WaitForServer(url); err != nil {
			log.Fatalf("Site is down: %s\n", err)
		}
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // 成功
		}
		log.SetPrefix("wait: ")
		log.SetFlags(0)
		log.Printf("server not responding(%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

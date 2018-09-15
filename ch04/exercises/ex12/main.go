package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Comic struct {
	Num        int
	Year       string
	Month      string
	Day        string
	Title      string
	Transcript string
	Alt        string
	Img        string
}

const (
	Max = 60
)

func geturl(num int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("Error")
	}

	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}

	return comic, nil
}

func getIdx(start, end int) ([]Comic, error) {
	var idx []Comic
	for i := start; i < end+1; i++ {
		c, err := geturl(i)
		if err != nil {
			return idx, err
		}
		idx = append(idx, c)
	}
	return idx, nil
}

var (
	s = flag.Int("s", 1, "start number")
	e = flag.Int("e", 1, "end number")
)

func main() {
	var idx []Comic
	var err error

	flag.Parse()
	if idx, err = getIdx(*s, *e); err != nil {
		log.Fatal("Error:", err)
	}
	for i, c := range idx {
		fmt.Printf("%d %v, %v, %v\n", i, c.Title, c.Transcript, c.Img)
	}

	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		txt := stdin.Text()
		for _, c := range idx {
			if strings.Index(c.Title, txt) != -1 || strings.Index(c.Transcript, txt) != -1 {
				fmt.Printf("Title: %s\nTranscript: %s\nImg: %s\n",
					c.Title, c.Transcript, c.Img)
			}
		}
	}
}

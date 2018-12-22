package main

import (
	"archive/tar"
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	format := flag.String("f", "zip", "unarchive format")
	flag.Parse()
	switch *format {
	case "zip":
		r, _ := zip.OpenReader(flag.Arg(0))
		defer r.Close()
		for _, f := range r.File {
			fmt.Println(f.Name)
		}
	case "tar":
		r, _ := os.Open(flag.Arg(0))
		defer r.Close()
		tr := tar.NewReader(r)
		for {
			h, err := tr.Next()
			if err == io.EOF {
				break
			}
			fmt.Println(h.Name)
		}
	}
}

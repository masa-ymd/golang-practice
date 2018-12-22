package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func pakages(patterns []string) []string {
	args := []string{"list", "-f={{.ImportPath}}"}
	for _, p := range patterns {
		args = append(args, p)
	}
	out, _ := exec.Command("go", args...).Output()
	targets := make(map[string]bool)
	for _, p := range strings.Fields(string(out)) {
		targets[p] = true
	}
	args2 := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	out2, _ := exec.Command("go", args2...).Output()
	var pkgs []string
	s := bufio.NewScanner(bytes.NewReader(out2))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		p := fields[0]
		d := fields[1:]
		for _, d := range d {
			if targets[d] {
				pkgs = append(pkgs, p)
				break
			}
		}
	}
	return pkgs
}

func main() {
	ps := pakages(os.Args[1:])
	for _, p := range ps {
		fmt.Println(p)
	}
}

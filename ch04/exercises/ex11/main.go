package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/masa-ymd/golang-practice/ch04/exercises/ex11/github"
)

var (
	o = flag.String("o", "search", "action")
	t = flag.String("t", "title", "title")
	n = flag.String("n", "number", "issue number")
	s = flag.String("s", "", "search string")
)

const (
	TmpFile = "tmpissuebody.txt"
)

func main() {
	flag.Parse()

	switch *o {
	case "search":
		sString := []string{"repo:masa-ymd/golang-practice"}
		sString = append(sString, *s)
		res, _ := github.SearchIssues(sString)
		fmt.Printf("%d issues:\n", res.TotalCount)
		for _, item := range res.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	case "edit":
		cmd := exec.Command("C:/Windows/notepad.exe", TmpFile)
		if err := cmd.Run(); err != nil {
			fmt.Errorf("exec error")
		}
		fp, err := os.Open(TmpFile)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			fp.Close()
			if err := os.Remove(TmpFile); err != nil {
				panic(err)
			}
		}()
		scanner := bufio.NewScanner(fp)
		var body string
		for scanner.Scan() {
			body = body + "\\n" + scanner.Text()
		}
		param := `{"title": "` + *t + `", "body": "` + body + `"}`
		if err := github.EditIssue(param, *n); err != nil {
			log.Fatal(err)
		}
	case "create":
		cmd := exec.Command("C:/Windows/notepad.exe", TmpFile)
		if err := cmd.Run(); err != nil {
			fmt.Errorf("exec error")
		}
		fp, err := os.Open(TmpFile)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			fp.Close()
			if err := os.Remove(TmpFile); err != nil {
				panic(err)
			}
		}()
		scanner := bufio.NewScanner(fp)
		var body string
		for scanner.Scan() {
			body = body + "\\n" + scanner.Text()
		}
		param := `{"title": "` + *t + `", "body": "` + body + `"}`
		if err := github.EditIssue(param, ""); err != nil {
			log.Fatal(err)
		}
	case "list":
		res, _ := github.SearchIssues([]string{"repo:masa-ymd/golang-practice"})
		fmt.Printf("%d issues:\n", res.TotalCount)
		for _, item := range res.Items {
			fmt.Printf("#%-5d %9.9s %.55s %s\n",
				item.Number, item.User.Login, item.Title, item.State)
		}
	case "open":
		param := `{"state": "open"}`
		if err := github.EditIssue(param, *n); err != nil {
			log.Fatal(err)
		}
	case "close":
		param := `{"state": "close"}`
		if err := github.EditIssue(param, *n); err != nil {
			log.Fatal(err)
		}
	}
}

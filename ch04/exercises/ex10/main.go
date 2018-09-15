package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/masa-ymd/golang-practice/ch04/exercises/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	afterM := time.Now()
	afterM = afterM.AddDate(0, -1, 0)
	afterY := time.Now()
	afterY = afterY.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		var c string
		switch {
		case item.CreatedAt.After(afterM):
			c = "Less than one month"
		case item.CreatedAt.Before(afterM) && item.CreatedAt.After(afterY):
			c = "More than one month"
		case item.CreatedAt.Before(afterY):
			c = "More than a year"
		}
		fmt.Printf("category:%s #%-5d %9.9s %.55s %s\n",
			c, item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

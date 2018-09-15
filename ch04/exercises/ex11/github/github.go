package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	SearchURL = "https://api.github.com/search/issues"
	RepoUrl   = "https://api.github.com/repos/masa-ymd/golang-practice/issues"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(SearchURL + "?q=" + q)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func EditIssue(param string, number string) error {
	var url string
	if number == "" {
		url = RepoUrl
	} else {
		url = strings.Join([]string{RepoUrl, number}, "/")
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(param)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	fmt.Print("User:")
	var gitUser string
	fmt.Scan(&gitUser)
	fmt.Print("password:")
	var gitPassword string
	fmt.Scan(&gitPassword)
	req.SetBasicAuth(gitUser, gitPassword)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

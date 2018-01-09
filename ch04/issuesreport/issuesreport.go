package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/masa-ymd/golang-practice/ch04/github"
)

const tmpl = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

// Mustはエラーがnilであればパニックを起こし、そうでなければテンプレートを返す
// Funcでテンプレート内の関数と自作関数の対応を定義
// パースで定数で定義したテンプレートを解析
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(tmpl))

func daysAgo(t time.Time) int {
	// Sinceでtを経過時間に変更
	return int(time.Since(t).Hours() / 24)
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

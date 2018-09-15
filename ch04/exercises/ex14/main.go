package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/masa-ymd/golang-practice/ch04/exercises/ex14/github"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, _ := github.SearchIssues([]string{"repo:masa-ymd/golang-practice"})
	t := template.Must(template.New("issuetmpl").Parse(`
	<!DOCTYPE html>
      <html>
        <body>
          {{ range . }}
			{{ $number := .Number }}
			{{ $user := .User.Login }}
			{{ $title := .Title }}
			{{ $state := .State }}
			{{ $milestone := .Milestone.Title }}
	        {{ range .Labels }}
	          {{ if eq .Name "bug" }}
                Number: {{ $number }}, User: {{ $user }} Title: {{ $title }} 
		        State: {{ $state }} Milestone {{ $milestone }} </br>
	          {{ end }}
            {{ end }}
	      {{ end }}
       </body>
    </html>
`))
	if err := t.ExecuteTemplate(w, "issuetmpl", res.Items); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

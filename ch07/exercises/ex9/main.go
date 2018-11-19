package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

type Track struct {
	Title string
	Year  int
}

type Tracks struct {
	T []Track
	C []condition
}

type comparison int

type condition func(a, b *Track) comparison

const (
	lt comparison = iota
	eq
	gt
)

func NewTracks(t []Track) *Tracks {
	return &Tracks{t, nil}
}

func (t *Tracks) LessTitle(a, b *Track) comparison {
	switch {
	case a.Title == b.Title:
		return eq
	case a.Title < b.Title:
		return lt
	default:
		return gt
	}
}

func (t *Tracks) LessYear(a, b *Track) comparison {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year < b.Year:
		return lt
	default:
		return gt
	}
}

func (t *Tracks) Len() int {
	return len(t.T)
}

func (t *Tracks) Swap(i, j int) {
	t.T[i], t.T[j] = t.T[j], t.T[i]
}

func (t *Tracks) Less(i, j int) bool {
	for _, f := range t.C {
		res := f(&t.T[i], &t.T[j])
		switch res {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

var trackList = template.Must(
	template.New("tracklist").Parse(`
	<html>
	<body>
	<table>
	    <tr>
		    <td><a href="?sort=title">Title</a></td>
			<td><a href="?sort=year">Year</a></td>
	    </tr>
	{{range .T}}
	    <tr>
	        <td>{{.Title}}</td>
			<td>{{.Year}}</td>
	    </tr>
	{{end}}
	</table>
	</body>
	</html>
	`),
)

func main() {
	tracks := []Track{
		{"A", 1},
		{"A", 2},
		{"B", 1},
		{"B", 2},
	}
	t := NewTracks(tracks)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(t.C) > 2 {
			t.C = t.C[1:]
		}
		switch r.FormValue("sort") {
		case "title":
			t.C = append(t.C, t.LessTitle)
		case "year":
			t.C = append(t.C, t.LessYear)
		}
		sort.Sort(t)
		if err := trackList.Execute(w, *t); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

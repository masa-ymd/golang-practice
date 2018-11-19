package main

import (
	"fmt"
	"sort"
)

type Track struct {
	Title string
	Year  int
}

type Tracks struct {
	t []Track
	c []condition
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
	return len(t.t)
}

func (t *Tracks) Swap(i, j int) {
	t.t[i], t.t[j] = t.t[j], t.t[i]
}

func (t *Tracks) Less(i, j int) bool {
	for _, f := range t.c {
		res := f(&t.t[i], &t.t[j])
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

func main() {
	tracks := []Track{
		{"A", 1},
		{"A", 2},
		{"B", 1},
		{"B", 2},
	}
	t := NewTracks(tracks)
	t.c = append(t.c, t.LessTitle)
	t.c = append(t.c, t.LessYear)
	sort.Sort(t)
	for i, x := range t.t {
		fmt.Printf("%d %v\n", i, x)
	}
	t.c = nil
	t.c = append(t.c, t.LessYear)
	t.c = append(t.c, t.LessTitle)
	sort.Sort(t)
	for i, x := range t.t {
		fmt.Printf("%d %v\n", i, x)
	}
}

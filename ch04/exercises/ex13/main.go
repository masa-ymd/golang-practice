package main

import (
	"encoding/json"
	"flag"
	//"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

type Result struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	RleaseDate       string  `json:"release_date"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float64 `json:"popularity"`
	VoteCount        int     `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
}

type Res struct {
	Page         int      `json:"page"`
	TotalResults int      `json:"total_results"`
	TotalPages   int      `json:"total_pages"`
	Results      []Result `json:"results"`
}

const (
	ApiURL = "https://api.themoviedb.org/3/search/movie?api_key="
)

var (
	apiKey = flag.String("k", "default", "apikey")
	query  = flag.String("q", "default", "Search String")
)

func main() {
	flag.Parse()
	url := ApiURL + *apiKey + "&query=" + *query
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("%v\n", err)
		os.Exit(1)
	}
	var r Res
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatalf("%v\n", err)
		os.Exit(1)
	}
	for _, t := range r.Results {
		if t.PosterPath == "" {
			continue
		}
		imgURL := "http://image.tmdb.org/t/p/w500" + t.PosterPath
		imgres, err := http.Get(imgURL)
		if err != nil {
			log.Fatalf("%v\n", err)
			os.Exit(1)
		}
		defer imgres.Body.Close()

		_, filename := path.Split(imgURL)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("%v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		io.Copy(file, imgres.Body)
	}
}

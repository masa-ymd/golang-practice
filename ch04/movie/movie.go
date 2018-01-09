package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"` // encoding/... の振る舞いを制御するフィールドタグ
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullit", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data2, err2 := json.MarshalIndent(movies, "", "    ")
	if err2 != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	// デコードする際は取得したい項目名を取得先の構造体に含める
	var titles []struct{ Title string }
	if err3 := json.Unmarshal(data, &titles); err3 != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err3)
	}
	fmt.Println(titles)
}

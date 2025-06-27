package main

import (
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func main() {
	tmdbClient, err := tmdb.Init(os.Getenv("APIKey"))

	if err != nil {
		fmt.Println(err)
	}

	options := make(map[string]string)
	options["language"] = "pt-BR"

	// Multi Search
	search, err := tmdbClient.GetSearchMulti("Dexter", options)

	if err != nil {
		fmt.Println(err)
	}

	// Iterate
	for _, result := range search.Results {
		switch media := result.(type) {
		case tmdb.MovieMedia:
			fmt.Println("Movie Title: ", media.Title)
		case tmdb.TVShowMedia:
			fmt.Println("TV Show: ", media.Name)
		case tmdb.PersonMedia:
			fmt.Println("Person: ", media.Name)
		}
	}
}

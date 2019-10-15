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
	for _, v := range search.Results {
		if v.MediaType == "movie" {
			fmt.Println("Movie Title: ", v.Title)
		} else if v.MediaType == "tv" {
			fmt.Println("TV Show: ", v.Name)
		} else if v.MediaType == "person" {
			fmt.Println("Person: ", v.Name)
		}
	}
}

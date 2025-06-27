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

	trending, err := tmdbClient.GetTrending("movie", "week", nil)

	if err != nil {
		fmt.Println(err)
	}

	for _, result := range trending.Results {
		switch media := result.(type) {
		case tmdb.MovieMedia:
			fmt.Println(media.Title)
		case tmdb.TVShowMedia:
			fmt.Println(media.OriginalName)
		}
	}

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["page"] = "2"
	options["language"] = "es-ES"

	trending, err = tmdbClient.GetTrending("tv", "day", options)

	if err != nil {
		fmt.Println(err)
	}

	for _, result := range trending.Results {
		switch media := result.(type) {
		case tmdb.MovieMedia:
			fmt.Println(media.Title)
		case tmdb.TVShowMedia:
			fmt.Println(media.OriginalName)
		}
	}
}

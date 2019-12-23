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

	if err := tmdbClient.SetSessionID(os.Getenv("SessionID")); err != nil {
		fmt.Println(err)
	}

	ratedMovies, err := tmdbClient.GetRatedMovies(0, nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, movie := range ratedMovies.Results {
		fmt.Println(movie.OriginalTitle)
		fmt.Println("")
	}
}

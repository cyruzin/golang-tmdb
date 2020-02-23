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

	// Enabling requests with context.
	tmdbClient.SetClientWithContext()

	options := map[string]string{
		"append_to_response": "videos",
	}

	movie, err := tmdbClient.GetMovieDetails(299536, options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Title)
}

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
	options["language"] = "en-US"
	options["page"] = "1"

	movie, err := tmdbClient.GetMovieUpcoming(options)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Results[0].Title)
	fmt.Println(movie.Results[0].GenreIDs)

}

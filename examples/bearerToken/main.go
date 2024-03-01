package main

import (
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func main() {
	tmdbClient, err := tmdb.InitV4(os.Getenv("BearerToken"))

	if err != nil {
		fmt.Println(err)
	}

	movie, err := tmdbClient.GetMovieDetails(299536, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Title)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tmdb "golang-tmdb"
)

func main() {

	tmdbClient, err := tmdb.Init(os.Getenv("APIKey"))

	tmdbClient.SetSessionID(os.Getenv("SessionID")) // From a ENV or Database

	// Setting a custom config for http.Client.
	tmdbClient.SetClientConfig(http.Client{Timeout: time.Second * 5})

	if err != nil {
		fmt.Println(err)
	}

	movie, err := tmdbClient.GetMovieDetails(299536, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Title)

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "credits"
	options["language"] = "pt-BR"

	movie, err = tmdbClient.GetMovieDetails(299536, options)

	if err != nil {
		fmt.Println(err)
	}

	// Credits - Iterate cast from append to response.
	for _, v := range movie.MovieCreditsAppend.Credits.Cast {
		fmt.Println(v.Name)
	}

	// PostMovieRating ..
	r, err := tmdbClient.PostMovieRating(299536, 3.5, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.StatusMessage)

	// DeleteMovieRating ..
	r, err = tmdbClient.DeleteMovieRating(299536, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.StatusMessage)
}

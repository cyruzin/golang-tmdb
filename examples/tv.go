package main

import (
	"fmt"
	"os"

	"github.com/cyruzin/golang-tmdb"
)

func main() {
	var tmdbClient tmdb.Client
	tmdbClient.APIKey = os.Getenv("APIKey")

	tvShow, err := tmdbClient.GetTVDetails(1399, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tvShow.Name)

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "credits"
	options["language"] = "pt-BR"

	tvShow, err = tmdbClient.GetTVDetails(1399, options)

	if err != nil {
		fmt.Println(err)
	}

	// Credits - Iterate cast from append to response.
	for _, v := range tvShow.TVCreditsAppend.Credits.Cast {
		fmt.Println(v.Name)
	}
}

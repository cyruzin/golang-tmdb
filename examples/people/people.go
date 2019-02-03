package main

import (
	"fmt"
	"os"

	"github.com/cyruzin/golang-tmdb"
)

func main() {
	var tmdbClient tmdb.Client
	tmdbClient.APIKey = os.Getenv("APIKey")

	people, err := tmdbClient.GetPeopleDetails(117642, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people.Name)

	fmt.Println("------")

	//With options
	options := make(map[string]string)
	options["append_to_response"] = "images"
	options["language"] = "pt-BR"

	people, err = tmdbClient.GetPeopleDetails(117642, options)

	if err != nil {
		fmt.Println(err)
	}

	// Images - Iterate profiles from append to response.
	for _, v := range people.PeopleImagesShort.Images.Profiles {
		fmt.Println(v.FilePath)
	}
}

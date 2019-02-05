package main

import (
	"fmt"
	"os"

	"github.com/cyruzin/golang-tmdb"
)

func main() {
	tmdbClient, err := tmdb.Init(os.Getenv("YOUR_APIKEY"))

	if err != nil {
		fmt.Println(err)
	}

	person, err := tmdbClient.GetPersonDetails(117642, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person.Name)

	fmt.Println("------")

	//With options
	options := make(map[string]string)
	options["append_to_response"] = "images"
	options["language"] = "pt-BR"

	person, err = tmdbClient.GetPersonDetails(117642, options)

	if err != nil {
		fmt.Println(err)
	}

	// Images - Iterate profiles from append to response.
	for _, v := range person.PersonImagesAppend.Images.Profiles {
		fmt.Println(v.FilePath)
	}
}

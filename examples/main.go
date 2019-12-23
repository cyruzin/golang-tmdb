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

	list := tmdb.ListCreate{
		Name:        "Test list!",
		Description: "Description is not enough...",
		Language:    "en",
	}

	r, err := tmdbClient.CreateList(&list)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.StatusCode)
	fmt.Println(r.StatusMessage)
}

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

	r, err := tmdbClient.DeleteList(128330)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r.StatusCode)
	fmt.Println(r.StatusMessage)
}

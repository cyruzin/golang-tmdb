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

	account, err := tmdbClient.GetAccountDetails()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.ID)
	fmt.Println(account.Username)
	fmt.Println(account.Name)
}

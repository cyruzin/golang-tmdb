// Package main demonstrates TMDB account API usage.
package main

import (
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func main() {
	tmdbClient, err := tmdb.New(
		tmdb.WithAPIKey(os.Getenv("APIKey")),
		tmdb.WithSessionID(os.Getenv("SessionID")),
	)
	if err != nil {
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

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

	options := map[string]string{
		"append_to_response": "videos",
	}

	movie, err := tmdbClient.GetMovieDetails(299536, options)
	if err != nil {
		fmt.Println(err)
	}

	// Generating Image URLs
	fmt.Println(tmdb.GetImageURL(movie.BackdropPath, "w500"))
	fmt.Println(tmdb.GetImageURL(movie.PosterPath, "original"))

	// Generating Video URLs
	for _, video := range movie.MovieVideosAppend.Videos.MovieVideos.Results {
		if video.Key != "" {
			fmt.Println(tmdb.GetVideoURL(video.Key))
		}
	}

}

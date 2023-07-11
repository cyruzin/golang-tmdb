package tmdb

import (
	"fmt"
)

// Trending type is a struct for trending JSON response.
type Trending struct {
	Page int `json:"page"`
	*TrendingResults
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetTrending get the daily or weekly trending items.
//
// The daily trending list tracks items over the period of
// a day while items have a 24 hour half life. The weekly
// list tracks items over a 7 day period, with a 7 day
// half life.
//
// # Valid Media Types
//
// all - Include all movies, TV shows and people in the
// results as a global trending list.
//
// movie - Show the trending movies in the results.
//
// tv - Show the trending TV shows in the results.
//
// person - Show the trending people in the results.
//
// # Valid Time Windows
//
// day - View the trending list for the day.
//
// week - View the trending list for the week.
//
// https://developers.themoviedb.org/3/trending/get-trending
func (c *Client) GetTrending(
	mediaType string,
	timeWindow string,
	options map[string]string,
) (*Trending, error) {
	tmdbURL := fmt.Sprintf(
		"%s/trending/%s/%s?api_key=%s",
		baseURL,
		mediaType,
		timeWindow,
		c.apiKey,
	)
	if options != nil {
		for key, value := range options {
			tmdbURL += fmt.Sprintf("&%s=%s", key, value)
		}
	}
	trending := Trending{}
	if err := c.get(tmdbURL, &trending); err != nil {
		return nil, err
	}
	return &trending, nil
}

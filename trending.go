package tmdb

import (
	"fmt"
)

// Trending type is a struct for treding JSON response.
type Trending struct {
	Page    int `json:"page"`
	Results []struct {
		Adult              bool     `json:"adult,omitempty"`
		Gender             int      `json:"gender,omitempty"`
		BackdropPath       string   `json:"backdrop_path,omitempty"`
		GenreIDs           []int64  `json:"genre_ids,omitempty"`
		ID                 int64    `json:"id"`
		OriginalLanguage   string   `json:"original_language"`
		OriginalTitle      string   `json:"original_title,omitempty"`
		Overview           string   `json:"overview,omitempty"`
		PosterPath         string   `json:"poster_path,omitempty"`
		ReleaseDate        string   `json:"release_date,omitempty"`
		Title              string   `json:"title,omitempty"`
		Video              bool     `json:"video,omitempty"`
		VoteAverage        float32  `json:"vote_average,omitempty"`
		VoteCount          int64    `json:"vote_count,omitempty"`
		Popularity         float32  `json:"popularity,omitempty"`
		FirstAirDate       string   `json:"first_air_date,omitempty"`
		Name               string   `json:"name,omitempty"`
		OriginCountry      []string `json:"origin_country,omitempty"`
		OriginalName       string   `json:"original_name,omitempty"`
		KnownForDepartment string   `json:"known_for_department,omitempty"`
		ProfilePath        string   `json:"profile_path,omitempty"`
		KnownFor           []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			GenreIds         []int   `json:"genre_ids"`
			ID               int     `json:"id"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date"`
			Title            string  `json:"title"`
			Video            bool    `json:"video"`
			VoteAverage      float64 `json:"vote_average"`
			VoteCount        int     `json:"vote_count"`
			Popularity       float64 `json:"popularity"`
			MediaType        string  `json:"media_type"`
		} `json:"known_for,omitempty"`
	} `json:"results"`
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
// Valid Media Types
//
// all - Include all movies, TV shows and people in the
// results as a global trending list.
//
// movie - Show the trending movies in the results.
//
// tv - Show the trending tv shows in the results.
//
// person - Show the trending people in the results.
//
// Valid Time Windows
//
// day - View the trending list for the day.
//
// week - View the trending list for the week.
//
// https://developers.themoviedb.org/3/trending/get-trending
func (c *Client) GetTrending(m, t string) (*Trending, error) {
	tmdbURL := fmt.Sprintf(
		"%s/trending/%s/%s?api_key=%s", baseURL, m, t, c.APIKey,
	)
	mt := Trending{}
	err := c.get(tmdbURL, &mt)
	if err != nil {
		return nil, err
	}
	return &mt, nil
}

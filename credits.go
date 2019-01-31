package tmdb

import (
	"encoding/json"
	"fmt"
)

// CreditsDetails type is a struct for credits JSON response.
type CreditsDetails struct {
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	Media      struct {
		OriginalName     string           `json:"original_name"`
		ID               int64            `json:"id"`
		Name             string           `json:"name"`
		VoteCount        int64            `json:"vote_count"`
		VoteAverage      float32          `json:"vote_average"`
		FirstAirDate     string           `json:"first_air_date"`
		PosterPath       string           `json:"poster_path"`
		GenreIDs         []int64          `json:"genre_ids"`
		OriginalLanguage string           `json:"original_language"`
		BackdropPath     string           `json:"backdrop_path"`
		Overview         string           `json:"overview"`
		OriginCountry    []string         `json:"origin_country"`
		Popularity       float32          `json:"popularity"`
		Character        string           `json:"character"`
		Episodes         *json.RawMessage `json:"episodes"`
		Seasons          []struct {
			AirDate      string `json:"air_date"`
			EpisodeCount int    `json:"episode_count"`
			ID           int64  `json:"id"`
			Name         string `json:"name"`
			Overview     string `json:"overview"`
			PosterPath   string `json:"poster_path"`
			SeasonNumber int    `json:"season_number"`
			ShowID       int64  `json:"show_id"`
		} `json:"seasons"`
	} `json:"media"`
	MediaType string `json:"media_type"`
	ID        string `json:"id"`
	Person    struct {
		Adult    bool   `json:"adult"`
		Gender   int    `json:"gender"`
		Name     string `json:"name"`
		ID       int64  `json:"id"`
		KnownFor []struct {
			Adult            bool     `json:"adult,omitempty"`
			BackdropPath     string   `json:"backdrop_path"`
			GenreIDs         []int64  `json:"genre_ids"`
			ID               int64    `json:"id"`
			OriginalLanguage string   `json:"original_language"`
			OriginalTitle    string   `json:"original_title,omitempty"`
			Overview         string   `json:"overview"`
			PosterPath       string   `json:"poster_path"`
			ReleaseDate      string   `json:"release_date,omitempty"`
			Title            string   `json:"title,omitempty"`
			Video            bool     `json:"video,omitempty"`
			VoteAverage      float32  `json:"vote_average"`
			VoteCount        int64    `json:"vote_count"`
			Popularity       float32  `json:"popularity"`
			MediaType        string   `json:"media_type"`
			OriginalName     string   `json:"original_name,omitempty"`
			Name             string   `json:"name,omitempty"`
			FirstAirDate     string   `json:"first_air_date,omitempty"`
			OriginCountry    []string `json:"origin_country,omitempty"`
		} `json:"known_for"`
		KnownForDepartment string  `json:"known_for_department"`
		ProfilePath        string  `json:"profile_path"`
		Popularity         float32 `json:"popularity"`
	} `json:"person"`
}

// GetCreditDetails get the list of timezones used throughout TMDb.
//
// https://developers.themoviedb.org/3/credits/get-credit-details
func (c *Client) GetCreditDetails(id string) (*CreditsDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%s?api_key=%s", baseURL, creditURL, id, c.APIKey,
	)
	t := CreditsDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

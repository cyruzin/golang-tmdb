package tmdb

import (
	"fmt"
)

// CreditsDetails type is a struct for credits JSON response.
type CreditsDetails struct {
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	Media      struct {
		Adult            bool     `json:"adult,omitempty"`          // Movie
		OriginalName     string   `json:"original_name,omitempty"`  // TV
		OriginalTitle    string   `json:"original_title,omitempty"` // Movie
		ID               int64    `json:"id"`
		Name             string   `json:"name,omitempty"` // TV
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		FirstAirDate     string   `json:"first_air_date,omitempty"` // TV
		PosterPath       string   `json:"poster_path"`
		ReleaseDate      string   `json:"release_date,omitempty"` // Movie
		Title            string   `json:"title,omitempty"`        // Movie
		Video            bool     `json:"video,omitempty"`        // Movie
		GenreIDs         []int64  `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country,omitempty"` // TV
		Popularity       float32  `json:"popularity"`
		Character        string   `json:"character"`
		Episodes         []struct {
			AirDate       string `json:"air_date"`
			EpisodeNumber int64  `json:"episode_number"`
			Name          string `json:"name"`
			Overview      string `json:"overview"`
			SeasonNumber  int    `json:"season_number"`
			StillPath     string `json:"still_path"`
		} `json:"episodes,omitempty"` // TV
		Seasons []Season `json:"seasons,omitempty"` // TV
	} `json:"media"`
	MediaType string `json:"media_type"`
	ID        string `json:"id"`
	Person    struct {
		Adult    bool   `json:"adult"`
		Gender   int    `json:"gender"`
		Name     string `json:"name"`
		ID       int64  `json:"id"`
		KnownFor []struct {
			Adult        bool   `json:"adult,omitempty"`
			BackdropPath string `json:"backdrop_path"`
			// GenreIDs         []int64  `json:"genre_ids"` // FIXME: -> []float32
			// ID               int64    `json:"id"` // FIXME: -> float32
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title,omitempty"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			ReleaseDate      string  `json:"release_date,omitempty"`
			Title            string  `json:"title,omitempty"`
			Video            bool    `json:"video,omitempty"`
			VoteAverage      float32 `json:"vote_average"`
			// VoteCount        int64    `json:"vote_count"` // FIXME: -> float32
			Popularity    float32  `json:"popularity"`
			MediaType     string   `json:"media_type"`
			OriginalName  string   `json:"original_name,omitempty"`
			Name          string   `json:"name,omitempty"`
			FirstAirDate  string   `json:"first_air_date,omitempty"`
			OriginCountry []string `json:"origin_country,omitempty"`
		} `json:"known_for"`
		KnownForDepartment string  `json:"known_for_department"`
		ProfilePath        string  `json:"profile_path"`
		Popularity         float32 `json:"popularity"`
	} `json:"person"`
}

// GetCreditDetails get a movie or TV credit details by id.
//
// https://developers.themoviedb.org/3/credits/get-credit-details
func (c *Client) GetCreditDetails(
	id string,
) (*CreditsDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%s?api_key=%s",
		baseURL,
		creditURL,
		id,
		c.apiKey,
	)
	creditsDetails := CreditsDetails{}
	if err := c.get(tmdbURL, &creditsDetails); err != nil {
		return nil, err
	}
	return &creditsDetails, nil
}

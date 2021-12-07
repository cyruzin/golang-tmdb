package tmdb

import "fmt"

// FindByID type is a struct for find JSON response.
type FindByID struct {
	MovieResults []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		VoteCount        int64   `json:"vote_count"`
		Popularity       float32 `json:"popularity"`
	} `json:"movie_results,omitempty"`
	PersonResults []struct {
		Adult    bool   `json:"adult"`
		Gender   int    `json:"gender"`
		Name     string `json:"name"`
		ID       int64  `json:"id"`
		KnownFor []struct {
			Adult        bool   `json:"adult,omitempty"` // Movie
			BackdropPath string `json:"backdrop_path"`
			FirstAirDate string `json:"first_air_date,omitempty"` // TV
			// GenreIDs         []int64 `json:"genre_ids"` // FIXME: -> []float32
			// ID               int64   `json:"id"` // FIXME: -> float32
			MediaType        string   `json:"media_type"`
			Name             string   `json:"name,omitempty"` // TV
			OriginalLanguage string   `json:"original_language"`
			OriginalName     string   `json:"original_name,omitempty"`  // TV
			OriginalTitle    string   `json:"original_title,omitempty"` // Movie
			OriginCountry    []string `json:"origin_country,omitempty"` // TV
			Overview         string   `json:"overview"`
			Popularity       float32  `json:"popularity"`
			PosterPath       string   `json:"poster_path"`
			ReleaseDate      string   `json:"release_date,omitempty"` // Movie
			Title            string   `json:"title,omitempty"`        // Movie
			Video            bool     `json:"video,omitempty"`        // Movie
			VoteAverage      float32  `json:"vote_average"`
			// VoteCount        int64   `json:"vote_count"` // FIXME: -> float32
		} `json:"known_for"`
		KnownForDepartment string  `json:"known_for_department"`
		ProfilePath        string  `json:"profile_path"`
		Popularity         float32 `json:"popularity"`
	} `json:"person_results,omitempty"`
	TvResults []struct {
		OriginalName     string   `json:"original_name"`
		ID               int64    `json:"id"`
		Name             string   `json:"name"`
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		FirstAirDate     string   `json:"first_air_date"`
		PosterPath       string   `json:"poster_path"`
		GenreIDs         []int64  `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		Popularity       float32  `json:"popularity"`
	} `json:"tv_results,omitempty"`
	TvEpisodeResults []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
	} `json:"tv_episode_results,omitempty"`
	TvSeasonResults []struct {
		AirDate      string `json:"air_date"`
		Name         string `json:"name"`
		ID           int64  `json:"id"`
		SeasonNumber int    `json:"season_number"`
		ShowID       int64  `json:"show_id"`
	} `json:"tv_season_results,omitempty"`
}

// GetFindByID the find method makes it easy to search for objects in our
// database by an external id. For example, an IMDB ID.
//
// This method will search all objects (movies, TV shows and people)
// and return the results in a single response.
//
// https://developers.themoviedb.org/3/find/find-by-id
func (c *Client) GetFindByID(
	id string,
	urlOptions map[string]string,
) (*FindByID, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s/find/%s?api_key=%s%s",
		baseURL, id, c.apiKey, options,
	)
	findByID := FindByID{}
	if err := c.get(tmdbURL, &findByID); err != nil {
		return nil, err
	}
	return &findByID, nil
}

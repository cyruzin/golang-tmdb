package tmdb

import (
	"fmt"
)

// Movies type is a struct for movies JSON response.
type Movies struct {
	Adult               bool   `json:"adult"`
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection"`
	Budget int64 `json:"budget"`
	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string  `json:"homepage"`
	ID                  int64   `json:"id"`
	IMDBID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"original_language"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float32 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_31661_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         int64  `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
}

// AlternativeTitles type is a struct for alternative titles JSON response.
type AlternativeTitles struct {
	ID     int `json:"id"`
	Titles []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string `json:"title"`
		Type      string `json:"type"`
	} `json:"titles"`
}

// GetMovieDetails get the primary information about a movie.
//
// https://developers.themoviedb.org/3/movies
//
func (c *Client) GetMovieDetails(id int, o map[string]string) (*Movies, error) {

	options := c.fmtOptions(o)

	tmdbURL := fmt.Sprintf("%s%s%d?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)

	var m Movies

	err := c.get(tmdbURL, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

// GetMovieAlternativeTitles get all of the alternative titles for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-alternative-titles
//
func (c *Client) GetMovieAlternativeTitles(id int, o map[string]string) (*AlternativeTitles, error) {

	options := c.fmtOptions(o)

	tmdbURL := fmt.Sprintf("%s%s%d/alternative_titles?api_key=%s%s", baseURL, movieURL, id, c.APIKey, options)

	var a AlternativeTitles

	err := c.get(tmdbURL, &a)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

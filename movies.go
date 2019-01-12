package tmdb

import "fmt"

// Movies type is a struct for movies JSON response.
type Movies struct {
	Adult        bool   `json:"adult"`
	BackdropPath string `json:"backdrop_path"`
	//BelongsToCollection
	Budget int64 `json:"budget"`
	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	Homepage            string  `json:"homepage"`
	ID                  int64   `json:"id"`
	IMDBID              string  `json:"imdb_id"`
	OriginalLanguage    string  `json:"string"`
	OriginalTitle       string  `json:"original_title"`
	Overview            string  `json:"overview"`
	Popularity          float32 `json:"popularity"`
	PosterPath          string  `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	}
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_31661_1"`
		Name      string `json:"name"`
	}
	ReleaseDate     string `json:"release_date"`
	Revenue         int64  `json:"revenue"`
	Runtime         int    `json:"runtime"`
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string `json:"name"`
	}
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
}

// GetDetails returns a movie details info.
//
// https://developers.themoviedb.org/3/movies
//
func (c *Client) GetDetails(id int) (*Movies, error) {
	tmdbURL := fmt.Sprintf("%s/movie/%d?api_key=%s", baseURL, id, c.APIKey)

	var m = &Movies{}

	err := c.get(tmdbURL, c)

	if err != nil {
		return nil, err
	}

	return m, nil
}

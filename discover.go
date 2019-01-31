package tmdb

import "fmt"

// DiscoverMovie type is a struct for movie JSON response.
type DiscoverMovie struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
		VoteCount        int64   `json:"vote_count"`
		ID               int64   `json:"id"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		GenreIDs         []int64 `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"results"`
}

// DiscoverTV type is a struct for tv JSON response.
type DiscoverTV struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
		OriginalName     string   `json:"original_name"`
		GenreIDs         []int64  `json:"genre_ids"`
		Name             string   `json:"name"`
		Popularity       float32  `json:"popularity"`
		OriginCountry    []string `json:"origin_country"`
		VoteCount        int64    `json:"vote_count"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		OriginalLanguage string   `json:"original_language"`
		ID               int64    `json:"id"`
		VoteAverage      float32  `json:"vote_average"`
		Overview         string   `json:"overview"`
		PosterPath       string   `json:"poster_path"`
	} `json:"results"`
}

// GetDiscoverMovie discover movies by different types of data like
// average rating, number of votes, genres and certifications. You can
// get a valid list of certifications from the  method.
//
// Discover also supports a nice list of sort options.
// See below for all of the available options.
//
// Please note, when using certification \ certification.lte you must
// also specify certification_country. These two parameters work together
// in order to filter the results. You can only filter results with the
// countries we have added to our certifications list.
//
// If you specify the region parameter, the regional release date will be
// used instead of the primary release date. The date returned will be the
// first date based on your query (ie. if a with_release_type is specified).
// It's important to note the order of the release types that are used.
// Specifying "2|3" would return the limited theatrical release date as
// opposed to "3|2" which would return the theatrical date.
//
// Also note that a number of filters support being comma (,) or pipe (|)
// separated. Comma's are treated like an AND and query while pipe's
// are an OR.
//
// https://developers.themoviedb.org/3/discover/movie-discover
func (c *Client) GetDiscoverMovie(o map[string]string) (*DiscoverMovie, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s%s", baseURL, discoverURL, c.APIKey, options,
	)
	t := DiscoverMovie{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetDiscoverTV Discover TV shows by different types of data like average
// rating, number of votes, genres, the network they aired on and air dates.
//
// Discover also supports a nice list of sort options. See below for all of
// the available options.
//
// Also note that a number of filters support being comma (,) or pipe (|)
// separated. Comma's are treated like an AND and query while pipe's are an OR.
//
// https://developers.themoviedb.org/3/discover/tv-discover
func (c *Client) GetDiscoverTV(o map[string]string) (*DiscoverTV, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s%s", baseURL, discoverURL, c.APIKey, options,
	)
	t := DiscoverTV{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

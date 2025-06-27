package tmdb

import "fmt"

// GetMovieNowPlaying get a list of movies in theatres.
//
// This is a release type query that looks for all movies that
// have a release type of 2 or 3 within the specified date range.
//
// You can optionally specify a region prameter which will narrow
// the search to only look for theatrical release dates within the
// specified country.
//
// https://developer.themoviedb.org/reference/movie-now-playing-list
func (c *Client) GetMovieNowPlaying(
	urlOptions map[string]string,
) (*PaginatedDatesMovieResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%snow_playing?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieNowPlaying := PaginatedDatesMovieResults{}
	if err := c.get(tmdbURL, &movieNowPlaying); err != nil {
		return nil, err
	}
	return &movieNowPlaying, nil
}

// GetMoviePopular get a list of the current popular movies on TMDb.
//
// This list updates daily.
//
// https://developer.themoviedb.org/reference/movie-popular-list
func (c *Client) GetMoviePopular(
	urlOptions map[string]string,
) (*PaginatedMovieResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	moviePopular := PaginatedMovieResults{}
	if err := c.get(tmdbURL, &moviePopular); err != nil {
		return nil, err
	}
	return &moviePopular, nil
}

// GetMovieTopRated get the top rated movies on TMDb.
//
// https://developer.themoviedb.org/reference/movie-top-rated-list
func (c *Client) GetMovieTopRated(
	urlOptions map[string]string,
) (*PaginatedMovieResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stop_rated?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieTopRated := PaginatedMovieResults{}
	if err := c.get(tmdbURL, &movieTopRated); err != nil {
		return nil, err
	}
	return &movieTopRated, nil
}

// GetMovieUpcoming get a list of upcoming movies in theatres.
//
// This is a release type query that looks for all movies that
// have a release type of 2 or 3 within the specified date range.
//
// You can optionally specify a region prameter which will narrow
// the search to only look for theatrical release dates within
// the specified country.
//
// https://developer.themoviedb.org/reference/movie-upcoming-list
func (c *Client) GetMovieUpcoming(
	urlOptions map[string]string,
) (*PaginatedDatesMovieResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%supcoming?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieUpcoming := PaginatedDatesMovieResults{}
	if err := c.get(tmdbURL, &movieUpcoming); err != nil {
		return nil, err
	}
	return &movieUpcoming, nil
}

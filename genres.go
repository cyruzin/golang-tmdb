package tmdb

import "fmt"

// GenresMovieList type is a struct for genres movie list JSON response.
type GenresMovieList struct {
	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}

// GetGenresMovieList get the list of official genres for movies.
//
// https://developers.themoviedb.org/3/keywords/get-keyword-details
func (c *Client) GetGenresMovieList(o map[string]string) (*GenresMovieList, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%smovie/list?api_key=%s%s", baseURL, genreURL, c.APIKey, options,
	)
	k := GenresMovieList{}
	err := c.get(tmdbURL, &k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// GetGenresTVList get the list of official genres for TV shows.
//
// https://developers.themoviedb.org/3/genres/get-tv-list
func (c *Client) GetGenresTVList(o map[string]string) (*GenresMovieList, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%stv/list?api_key=%s%s", baseURL, genreURL, c.APIKey, options,
	)
	k := GenresMovieList{}
	err := c.get(tmdbURL, &k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

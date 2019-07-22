package tmdb

import "fmt"

// GenreMovieList type is a struct for genres movie list JSON response.
type GenreMovieList struct {
	Genres []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}

// GetGenreMovieList get the list of official genres for movies.
//
// https://developers.themoviedb.org/3/keywords/get-keyword-details
func (c *Client) GetGenreMovieList(
	o map[string]string,
) (*GenreMovieList, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%smovie/list?api_key=%s%s",
		baseURL, genreURL, c.apiKey, options,
	)
	k := GenreMovieList{}
	err := c.get(tmdbURL, &k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// GetGenreTVList get the list of official genres for TV shows.
//
// https://developers.themoviedb.org/3/genres/get-tv-list
func (c *Client) GetGenreTVList(
	o map[string]string,
) (*GenreMovieList, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%stv/list?api_key=%s%s",
		baseURL, genreURL, c.apiKey, options,
	)
	k := GenreMovieList{}
	err := c.get(tmdbURL, &k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

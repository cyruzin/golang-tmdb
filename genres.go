package tmdb

import "fmt"

// GenreMovieList type is a struct for genres movie list JSON response.
type GenreMovieList struct {
	Genres []Genre `json:"genres"`
}

// GetGenreMovieList get the list of official genres for movies.
//
// https://developers.themoviedb.org/3/genres/get-movie-list
func (c *Client) GetGenreMovieList(
	urlOptions map[string]string,
) (*GenreMovieList, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie/list?api_key=%s%s",
		baseURL,
		genreURL,
		c.apiKey,
		options,
	)
	genreMovieList := GenreMovieList{}
	if err := c.get(tmdbURL, &genreMovieList); err != nil {
		return nil, err
	}
	return &genreMovieList, nil
}

// GetGenreTVList get the list of official genres for TV shows.
//
// https://developers.themoviedb.org/3/genres/get-tv-list
func (c *Client) GetGenreTVList(
	urlOptions map[string]string,
) (*GenreMovieList, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv/list?api_key=%s%s",
		baseURL,
		genreURL,
		c.apiKey,
		options,
	)
	genreTVList := GenreMovieList{}
	if err := c.get(tmdbURL, &genreTVList); err != nil {
		return nil, err
	}
	return &genreTVList, nil
}

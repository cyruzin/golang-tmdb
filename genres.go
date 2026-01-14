package tmdb

// GenreMovieList type is a struct for genres movie list JSON response.
type GenreMovieList struct {
	Genres []Genre `json:"genres"`
}

// GetGenreMovieList get the list of official genres for movies.
//
// https://developers.themoviedb.org/3/genres/get-movie-list
func (c *Client) GetGenreMovieList(urlOptions map[string]string) (*GenreMovieList, error) {
	url := c.baseURL.JoinPath(genreURL, "movie", "list")
	var genreMovieList GenreMovieList
	if err := c.get(url, urlOptions, &genreMovieList); err != nil {
		return nil, err
	}
	return &genreMovieList, nil
}

// GetGenreTVList get the list of official genres for TV shows.
//
// https://developers.themoviedb.org/3/genres/get-tv-list
func (c *Client) GetGenreTVList(urlOptions map[string]string) (*GenreMovieList, error) {
	url := c.baseURL.JoinPath(genreURL, "tv", "list")
	var genreTVList GenreMovieList
	if err := c.get(url, urlOptions, &genreTVList); err != nil {
		return nil, err
	}
	return &genreTVList, nil
}

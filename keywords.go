package tmdb

import "fmt"

// GetKeywordDetails get keyword details by id.
//
// https://developer.themoviedb.org/reference/keyword-details
func (c *Client) GetKeywordDetails(
	id int,
) (*KeywordDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s",
		baseURL,
		keywordURL,
		id,
		c.apiKey,
	)
	keywordDetails := KeywordDetails{}
	if err := c.get(tmdbURL, &keywordDetails); err != nil {
		return nil, err
	}
	return &keywordDetails, nil
}

// Deprecated: Use GetDiscoverMovie instead.
//
// GetKeywordMovies get the movies that belong to a keyword.
//
// We highly recommend using movie discover instead of this
// method as it is much more flexible.
//
// This endpoint is deprecated in the TMDB API.
//
// https://developer.themoviedb.org/reference/keyword-movies
func (c *Client) GetKeywordMovies(
	id int,
	urlOptions map[string]string,
) (*KeywordMovies, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/movies?api_key=%s%s",
		baseURL,
		keywordURL,
		id,
		c.apiKey,
		options,
	)
	keywordMovies := KeywordMovies{}
	if err := c.get(tmdbURL, &keywordMovies); err != nil {
		return nil, err
	}
	return &keywordMovies, nil
}

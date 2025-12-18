package tmdb

import "fmt"

// GetChangesMovie get a list of all of the movie ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developer.themoviedb.org/reference/changes-movie-list
func (c *Client) GetChangesMovie(
	urlOptions map[string]string,
) (*ChangesMovie, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%schanges?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	changesMovies := ChangesMovie{}
	if err := c.get(tmdbURL, &changesMovies); err != nil {
		return nil, err
	}
	return &changesMovies, nil
}

// GetChangesPerson get a list of all of the person ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developer.themoviedb.org/reference/changes-people-list
func (c *Client) GetChangesPerson(
	urlOptions map[string]string,
) (*ChangesPerson, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%schanges?api_key=%s%s",
		baseURL,
		personURL,
		c.apiKey,
		options,
	)
	changesPerson := ChangesPerson{}
	if err := c.get(tmdbURL, &changesPerson); err != nil {
		return nil, err
	}
	return &changesPerson, nil
}

// GetChangesTV get a list of all of the TV show ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developer.themoviedb.org/reference/changes-tv-list
func (c *Client) GetChangesTV(
	urlOptions map[string]string,
) (*ChangesTV, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%schanges?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	changesTV := ChangesTV{}
	if err := c.get(tmdbURL, &changesTV); err != nil {
		return nil, err
	}
	return &changesTV, nil
}

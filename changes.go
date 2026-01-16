package tmdb

// ChangesMovie type is a struct for movie changes JSON response.
type ChangesMovie struct {
	*ChangesMovieResults
	PaginatedResultsMeta
}

// GetChangesMovie get a list of all of the movie ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-movie-change-list
func (c *Client) GetChangesMovie(urlOptions map[string]string) (*ChangesMovie, error) {
	url := c.baseURL.JoinPath(moviePath, "changes")
	var changesMovies ChangesMovie
	if err := c.get(url, urlOptions, &changesMovies); err != nil {
		return nil, err
	}
	return &changesMovies, nil
}

// ChangesTV type is a struct for tv changes JSON response.
type ChangesTV struct {
	*ChangesMovie
}

// GetChangesTV get a list of all of the TV show ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-tv-change-list
func (c *Client) GetChangesTV(urlOptions map[string]string) (*ChangesTV, error) {
	url := c.baseURL.JoinPath(tvPath, "changes")
	var changesTV ChangesTV
	if err := c.get(url, urlOptions, &changesTV); err != nil {
		return nil, err
	}
	return &changesTV, nil
}

// ChangesPerson type is a struct for person changes JSON response.
type ChangesPerson struct {
	*ChangesMovie
}

// GetChangesPerson get a list of all of the person ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-person-change-list
func (c *Client) GetChangesPerson(urlOptions map[string]string) (*ChangesPerson, error) {
	url := c.baseURL.JoinPath(personPath, "changes")
	var changesPerson ChangesPerson
	if err := c.get(url, urlOptions, &changesPerson); err != nil {
		return nil, err
	}
	return &changesPerson, nil
}

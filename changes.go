package tmdb

import "fmt"

// ChangesMovie type is a struct for movie changes JSON response.
type ChangesMovie struct {
	Results []struct {
		ID    int64 `json:"id"`
		Adult bool  `json:"adult"`
	} `json:"results"`
	Page         int64 `json:"page"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// ChangesTV type is a struct for tv changes JSON response.
type ChangesTV struct {
	*ChangesMovie
}

// ChangesPerson type is a struct for person changes JSON response.
type ChangesPerson struct {
	*ChangesMovie
}

// GetChangesMovie get a list of all of the movie ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-movie-change-list
func (c *Client) GetChangesMovie(o map[string]string) (*ChangesMovie, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%schanges?api_key=%s%s", baseURL, movieURL, c.APIKey, options)
	m := ChangesMovie{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetChangesTV get a list of all of the TV show ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-tv-change-list
func (c *Client) GetChangesTV(o map[string]string) (*ChangesTV, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%schanges?api_key=%s%s", baseURL, tvURL, c.APIKey, options)
	m := ChangesTV{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetChangesPerson get a list of all of the person ids
// that have been changed in the past 24 hours.
//
// You can query it for up to 14 days worth of changed IDs
// at a time with the start_date and end_date query parameters.
// 100 items are returned per page.
//
// https://developers.themoviedb.org/3/changes/get-person-change-list
func (c *Client) GetChangesPerson(o map[string]string) (*ChangesPerson, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%schanges?api_key=%s%s", baseURL, personURL, c.APIKey, options)
	m := ChangesPerson{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

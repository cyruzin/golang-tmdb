package tmdb

import (
	"fmt"
	"net/url"
)

// GetSearchCompanies search for companies.
//
// https://developer.themoviedb.org/reference/search-company
func (c *Client) GetSearchCompanies(
	query string,
	urlOptions map[string]string,
) (*SearchCompanies, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%scompany?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	SearchCompanies := SearchCompanies{}
	if err := c.get(tmdbURL, &SearchCompanies); err != nil {
		return nil, err
	}
	return &SearchCompanies, nil
}

// GetSearchCollections search for collections.
//
// https://developer.themoviedb.org/reference/search-collection
func (c *Client) GetSearchCollections(
	query string,
	urlOptions map[string]string,
) (*SearchCollections, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%scollection?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchCollections := SearchCollections{}
	if err := c.get(tmdbURL, &searchCollections); err != nil {
		return nil, err
	}
	return &searchCollections, nil
}

// GetSearchKeywords search for keywords.
//
// https://developer.themoviedb.org/reference/search-keyword
func (c *Client) GetSearchKeywords(
	query string,
	urlOptions map[string]string,
) (*SearchKeywords, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%skeyword?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchKeywords := SearchKeywords{}
	if err := c.get(tmdbURL, &searchKeywords); err != nil {
		return nil, err
	}
	return &searchKeywords, nil
}

// GetSearchMovies search for keywords.
//
// https://developer.themoviedb.org/reference/search-movie
func (c *Client) GetSearchMovies(
	query string,
	urlOptions map[string]string,
) (*SearchMovies, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchMovies := SearchMovies{}
	if err := c.get(tmdbURL, &searchMovies); err != nil {
		return nil, err
	}
	return &searchMovies, nil
}

// GetSearchMulti search multiple models in a single request.
// Multi search currently supports searching for movies,
// tv shows and people in a single request.
//
// https://developer.themoviedb.org/reference/search-multi
func (c *Client) GetSearchMulti(
	query string,
	urlOptions map[string]string,
) (*SearchMulti, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smulti?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchMulti := SearchMulti{}
	if err := c.get(tmdbURL, &searchMulti); err != nil {
		return nil, err
	}
	return &searchMulti, nil
}

// GetSearchPeople search for people.
//
// https://developer.themoviedb.org/reference/search-person
func (c *Client) GetSearchPeople(
	query string,
	urlOptions map[string]string,
) (*SearchPeople, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sperson?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchPeople := SearchPeople{}
	if err := c.get(tmdbURL, &searchPeople); err != nil {
		return nil, err
	}
	return &searchPeople, nil
}

// GetSearchTVShow search for a TV Show.
//
// https://developer.themoviedb.org/reference/search-tv
func (c *Client) GetSearchTVShow(
	query string,
	urlOptions map[string]string,
) (*SearchTVShows, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchTVShows := SearchTVShows{}
	if err := c.get(tmdbURL, &searchTVShows); err != nil {
		return nil, err
	}
	return &searchTVShows, nil
}

package tmdb

import (
	"fmt"
	"net/url"
)

// GetSearchCompanies search for companies.
//
// https://developers.themoviedb.org/3/search/search-companies
func (c *Client) GetSearchCompanies(
	query string,
	urlOptions map[string]string,
) (*PaginatedCompanyInfoResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%scompany?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	SearchCompanies := PaginatedCompanyInfoResults{}
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
) (*PaginatedCollectionResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%scollection?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchCollections := PaginatedCollectionResults{}
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
) (*PaginatedKeywordsResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%skeyword?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchKeywords := PaginatedKeywordsResults{}
	if err := c.get(tmdbURL, &searchKeywords); err != nil {
		return nil, err
	}
	return &searchKeywords, nil
}

// GetSearchMovies search for keywords.
//
// https://developers.themoviedb.org/3/search/search-movies
func (c *Client) GetSearchMovies(
	query string,
	urlOptions map[string]string,
) (*PaginatedMovieMediaResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchMovies := PaginatedMovieMediaResults{}
	if err := c.get(tmdbURL, &searchMovies); err != nil {
		return nil, err
	}
	return &searchMovies, nil
}

// GetSearchMulti search multiple models in a single request.
// Multi search currently supports searching for movies,
// tv shows and people in a single request.
//
// https://developers.themoviedb.org/3/search/multi-search
func (c *Client) GetSearchMulti(
	query string,
	urlOptions map[string]string,
) (*PaginatedMediaResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smulti?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchMulti := PaginatedMediaResults{}
	if err := c.get(tmdbURL, &searchMulti); err != nil {
		return nil, err
	}
	return &searchMulti, nil
}

// GetSearchPeople search for people.
//
// https://developers.themoviedb.org/3/search/search-people
func (c *Client) GetSearchPeople(
	query string,
	urlOptions map[string]string,
) (*PaginatedPersonResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sperson?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchPeople := PaginatedPersonResults{}
	if err := c.get(tmdbURL, &searchPeople); err != nil {
		return nil, err
	}
	return &searchPeople, nil
}

// GetSearchTVShow search for a TV Show.
//
// https://developers.themoviedb.org/3/search/search-tv-shows
func (c *Client) GetSearchTVShow(
	query string,
	urlOptions map[string]string,
) (*PaginatedTVShowResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s&query=%s%s",
		baseURL,
		searchURL,
		c.apiKey,
		url.QueryEscape(query),
		options,
	)
	searchTVShows := PaginatedTVShowResults{}
	if err := c.get(tmdbURL, &searchTVShows); err != nil {
		return nil, err
	}
	return &searchTVShows, nil
}

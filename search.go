package tmdb

import (
	"fmt"
	"net/url"
)

// SearchCompanies type is a struct for companies JSON response.
type SearchCompanies struct {
	*SearchCompaniesResults
	PaginatedResultsMeta
}

// GetSearchCompanies search for companies.
//
// https://developers.themoviedb.org/3/search/search-companies
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

// SearchCollections type is a strcut for collections JSON response.
type SearchCollections struct {
	*SearchCollectionsResults
	PaginatedResultsMeta
}

// GetSearchCollections search for collections.
//
// https://developers.themoviedb.org/3/search/search-collections
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

// SearchKeywords type is a struct for keywords JSON response.
type SearchKeywords struct {
	*SearchKeywordsResults
	PaginatedResultsMeta
}

// GetSearchKeywords search for keywords.
//
// https://developers.themoviedb.org/3/search/search-keywords
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

// SearchMovies type is a struct for movies JSON response.
type SearchMovies struct {
	PaginatedResultsMeta
	*SearchMoviesResults
}

// GetSearchMovies search for keywords.
//
// https://developers.themoviedb.org/3/search/search-movies
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

// SearchMulti type is a struct for multi JSON response.
type SearchMulti struct {
	*SearchMultiResults
	PaginatedResultsMeta
}

// GetSearchMulti search multiple models in a single request.
// Multi search currently supports searching for movies,
// tv shows and people in a single request.
//
// https://developers.themoviedb.org/3/search/multi-search
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

// SearchPeople type is a struct for people JSON response.
type SearchPeople struct {
	PaginatedResultsMeta
	*SearchPeopleResults
}

// GetSearchPeople search for people.
//
// https://developers.themoviedb.org/3/search/search-people
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

// SearchTVShows type is a struct for tv show JSON response.
type SearchTVShows struct {
	PaginatedResultsMeta
	*SearchTVShowsResults
}

// GetSearchTVShow search for a TV Show.
//
// https://developers.themoviedb.org/3/search/search-tv-shows
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

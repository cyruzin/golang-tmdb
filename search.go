package tmdb

import (
	"fmt"
	"net/url"
)

// SearchCompanies type is a struct for companies JSON response.
type SearchCompanies struct {
	Page    int64 `json:"page"`
	Results []struct {
		ID       int64  `json:"id"`
		LogoPath string `json:"logo_path"`
		Name     string `json:"name"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// SearchCollections type is a strcut for collections JSON response.
type SearchCollections struct {
	Page    int64 `json:"page"`
	Results []struct {
		Adult            bool   `json:"adult"`
		BackdropPath     string `json:"backdrop_path"`
		ID               int64  `json:"id"`
		Name             string `json:"name"`
		OriginalLanguage string `json:"original_language"`
		OriginalName     string `json:"original_name"`
		Overview         string `json:"overview"`
		PosterPath       string `json:"poster_path"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// SearchKeywords type is a struct for keywords JSON response.
type SearchKeywords struct {
	Page    int64 `json:"page"`
	Results []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// SearchMovies type is a struct for movies JSON response.
type SearchMovies struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
		VoteCount        int64   `json:"vote_count"`
		ID               int64   `json:"id"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		GenreIds         []int   `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"results"`
}

// GetSearchCompanies search for companies.
//
// https://developers.themoviedb.org/3/search/search-companies
func (c *Client) GetSearchCompanies(q string, o map[string]string) (*SearchCompanies, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%scompany?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchCompanies{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetSearchCollections search for collections.
//
// https://developers.themoviedb.org/3/search/search-collections
func (c *Client) GetSearchCollections(q string, o map[string]string) (*SearchCollections, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%scollection?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchCollections{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetSearchKeywords search for keywords.
//
// https://developers.themoviedb.org/3/search/search-keywords
func (c *Client) GetSearchKeywords(q string, o map[string]string) (*SearchKeywords, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%skeyword?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchKeywords{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetSearchMovies search for keywords.
//
// https://developers.themoviedb.org/3/search/search-movies
func (c *Client) GetSearchMovies(q string, o map[string]string) (*SearchMovies, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchMovies{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

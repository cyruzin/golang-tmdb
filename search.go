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
		GenreIDs         []int64 `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"results"`
}

// SearchMulti type is a struct for multi JSON response.
type SearchMulti struct {
	Page    int `json:"page"`
	Results []struct {
		PosterPath       string   `json:"poster_path,omitempty"`
		Popularity       float32  `json:"popularity"`
		ID               int64    `json:"id"`
		Overview         string   `json:"overview,omitempty"`
		BackdropPath     string   `json:"backdrop_path,omitempty"`
		VoteAverage      float32  `json:"vote_average,omitempty"`
		MediaType        string   `json:"media_type"`
		FirstAirDate     string   `json:"first_air_date,omitempty"`
		OriginCountry    []string `json:"origin_country,omitempty"`
		GenreIDs         []int64  `json:"genre_ids,omitempty"`
		OriginalLanguage string   `json:"original_language,omitempty"`
		VoteCount        int64    `json:"vote_count,omitempty"`
		Name             string   `json:"name,omitempty"`
		OriginalName     string   `json:"original_name,omitempty"`
		Adult            bool     `json:"adult,omitempty"`
		ReleaseDate      string   `json:"release_date,omitempty"`
		OriginalTitle    string   `json:"original_title,omitempty"`
		Title            string   `json:"title,omitempty"`
		Video            bool     `json:"video,omitempty"`
		ProfilePath      string   `json:"profile_path,omitempty"`
		KnownFor         []struct {
			PosterPath       string  `json:"poster_path"`
			Adult            bool    `json:"adult"`
			Overview         string  `json:"overview"`
			ReleaseDate      string  `json:"release_date"`
			OriginalTitle    string  `json:"original_title"`
			GenreIDs         []int64 `json:"genre_ids"`
			ID               int64   `json:"id"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			Title            string  `json:"title"`
			BackdropPath     string  `json:"backdrop_path"`
			Popularity       float32 `json:"popularity"`
			VoteCount        int64   `json:"vote_count"`
			Video            bool    `json:"video"`
			VoteAverage      float32 `json:"vote_average"`
		} `json:"known_for,omitempty"`
	} `json:"results"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
}

// SearchPeople type is a struct for people JSON response.
type SearchPeople struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
		Popularity  float32 `json:"popularity"`
		ID          int64   `json:"id"`
		ProfilePath string  `json:"profile_path"`
		Name        string  `json:"name"`
		KnownFor    []struct {
			VoteAverage      float32 `json:"vote_average"`
			VoteCount        int64   `json:"vote_count"`
			ID               int64   `json:"id"`
			Video            bool    `json:"video"`
			MediaType        string  `json:"media_type"`
			Title            string  `json:"title"`
			Popularity       float32 `json:"popularity"`
			PosterPath       string  `json:"poster_path"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			GenreIDs         []int64 `json:"genre_ids"`
			BackdropPath     string  `json:"backdrop_path"`
			Adult            bool    `json:"adult"`
			Overview         string  `json:"overview"`
			ReleaseDate      string  `json:"release_date"`
		} `json:"known_for"`
		Adult bool `json:"adult"`
	} `json:"results"`
}

// SearchTVShows type is a struct for tv show JSON response.
type SearchTVShows struct {
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
		OriginalName     string   `json:"original_name"`
		ID               int64    `json:"id"`
		Name             string   `json:"name"`
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		Popularity       float32  `json:"popularity"`
		GenreIDs         []int64  `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
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

// GetSearchMulti search multiple models in a single request.
// Multi search currently supports searching for movies,
// tv shows and people in a single request.
//
// https://developers.themoviedb.org/3/search/multi-search
func (c *Client) GetSearchMulti(q string, o map[string]string) (*SearchMulti, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%smulti?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchMulti{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetSearchPeople search for people.
//
// https://developers.themoviedb.org/3/search/search-people
func (c *Client) GetSearchPeople(q string, o map[string]string) (*SearchPeople, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%sperson?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchPeople{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// GetSearchTVShow search for a TV Show.
//
// https://developers.themoviedb.org/3/search/search-tv-shows
func (c *Client) GetSearchTVShow(q string, o map[string]string) (*SearchTVShows, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s&query=%s%s", baseURL, searchURL, c.APIKey, url.QueryEscape(q), options,
	)
	s := SearchTVShows{}
	err := c.get(tmdbURL, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

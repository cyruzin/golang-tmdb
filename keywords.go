package tmdb

import "fmt"

// KeywordDetails type is a struct for keyword JSON response.
type KeywordDetails struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetKeywordDetails get keyword details by id.
//
// https://developers.themoviedb.org/3/keywords/get-keyword-details
func (c *Client) GetKeywordDetails(id int) (*KeywordDetails, error) {
	url := c.baseURL.JoinPath(keywordPath, fmt.Sprintf("%d", id))
	var keywordDetails KeywordDetails
	if err := c.get(url, nil, &keywordDetails); err != nil {
		return nil, err
	}
	return &keywordDetails, nil
}

// KeywordMovies type is a struct for movies that belong to a keyword JSON response.
type KeywordMovies struct {
	ID      int `json:"id"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		Popularity       float32 `json:"popularity"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

// GetKeywordMovies get the movies that belong to a keyword.
//
// We highly recommend using movie discover instead of this
// method as it is much more flexible.
//
// https://developers.themoviedb.org/3/keywords/get-movies-by-keyword
func (c *Client) GetKeywordMovies(id int, urlOptions map[string]string) (*KeywordMovies, error) {
	url := c.baseURL.JoinPath(keywordPath, fmt.Sprintf("%d", id), "movies")
	var keywordMovies KeywordMovies
	if err := c.get(url, urlOptions, &keywordMovies); err != nil {
		return nil, err
	}
	return &keywordMovies, nil
}

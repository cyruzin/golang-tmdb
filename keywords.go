package tmdb

import "fmt"

// KeywordDetails type is a struct for keyword JSON response.
type KeywordDetails struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// KeywordMovies type is a struct for movies that belong to a keyword JSON response.
type KeywordMovies struct {
	ID      int64 `json:"id"`
	Page    int64 `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		VoteCount        int64   `json:"vote_count"`
		Popularity       float32 `json:"popularity"`
	} `json:"results"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
}

// GetKeywordDetails get keyword details by id.
//
// https://developers.themoviedb.org/3/keywords/get-keyword-details
func (c *Client) GetKeywordDetails(id int) (*KeywordDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s",
		baseURL, keywordURL, id, c.APIKey,
	)
	k := KeywordDetails{}
	err := c.get(tmdbURL, &k)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// GetKeywordMovies get the movies that belong to a keyword.
//
// We highly recommend using movie discover instead of this
// method as it is much more flexible.
//
// https://developers.themoviedb.org/3/keywords/get-movies-by-keyword
func (c *Client) GetKeywordMovies(
	id int, o map[string]string,
) (*KeywordMovies, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/movies?api_key=%s%s",
		baseURL, keywordURL, id, c.APIKey, options,
	)
	t := KeywordMovies{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

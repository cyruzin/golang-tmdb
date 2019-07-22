package tmdb

import "fmt"

// ListDetails type is a struct for details JSON response.
type ListDetails struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            string `json:"id"`
	Items         []struct {
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
	} `json:"items"`
	ItemCount  int64  `json:"item_count"`
	Iso639_1   string `json:"iso_639_1"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
}

// ListItemStatus type is a struct for item status JSON response.
type ListItemStatus struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// GetListDetails get the details of a list.
//
// https://developers.themoviedb.org/3/lists/get-list-details
func (c *Client) GetListDetails(
	id string, o map[string]string,
) (*ListDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%s?api_key=%s%s",
		baseURL, listURL, id, c.apiKey, options,
	)
	t := ListDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetListItemStatus check if a movie has already been added to the list.
//
// https://developers.themoviedb.org/3/lists/check-item-status
func (c *Client) GetListItemStatus(
	id string, o map[string]string,
) (*ListItemStatus, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%s/item_status?api_key=%s%s",
		baseURL, listURL, id, c.apiKey, options,
	)
	t := ListItemStatus{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

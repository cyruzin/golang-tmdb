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

// ListResponse type is a struct for list creation JSON response.
type ListResponse struct {
	*Response
	Success bool  `json:"success"`
	ListID  int64 `json:"list_id"`
}

// ListCreate type is a struct for list creation JSON request.
type ListCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

// ListMedia type is a struct for list media JSON request.
type ListMedia struct {
	MediaID int64 `json:"media_id"`
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

// CreateList creates a list.
//
// https://developers.themoviedb.org/3/lists/create-list
func (c *Client) CreateList(list *ListCreate) (*ListResponse, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list?api_key=%s&session_id=%s",
		baseURL, c.apiKey, c.sessionID,
	)
	createList := ListResponse{}
	err := c.request(tmdbURL, list, "POST", &createList)
	if err != nil {
		return nil, err
	}
	return &createList, nil
}

// ClearList clear all of the items from a list.
//
// https://developers.themoviedb.org/3/lists/clear-list
func (c *Client) ClearList(listID int, confirm bool) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d/clear?api_key=%s&session_id=%s&confirm=%t",
		baseURL, listID, c.apiKey, c.sessionID, confirm,
	)
	clearList := Response{}
	err := c.request(tmdbURL, listID, "POST", &clearList)
	if err != nil {
		return nil, err
	}
	return &clearList, nil
}

// DeleteList deletes a list.
//
// https://developers.themoviedb.org/3/lists/delete-list
func (c *Client) DeleteList(listID int) (*Response, error) {
	//https://api.themoviedb.org/3/list/12121?api_key=jsasaisajs&session_id=asasas
	tmdbURL := fmt.Sprintf(
		"%s/list/%d?api_key=%s&session_id=%s",
		baseURL, listID, c.apiKey, c.sessionID,
	)
	deleteList := Response{}
	err := c.request(tmdbURL, nil, "DELETE", &deleteList)
	if err != nil {
		return nil, err
	}
	return &deleteList, nil
}

// AddMovie add a movie to a list.
//
// https://developers.themoviedb.org/3/lists/add-movie
func (c *Client) AddMovie(listID int, mediaID *ListMedia) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d/add_item?api_key=%s&session_id=%s",
		baseURL, listID, c.apiKey, c.sessionID,
	)
	addMovie := Response{}
	err := c.request(tmdbURL, mediaID, "POST", &addMovie)
	if err != nil {
		return nil, err
	}
	return &addMovie, nil
}

// RemoveMovie remove a movie from a list.
//
// https://developers.themoviedb.org/3/lists/remove-movie
func (c *Client) RemoveMovie(listID int, mediaID *ListMedia) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d/remove_item?api_key=%s&session_id=%s",
		baseURL, listID, c.apiKey, c.sessionID,
	)
	removeMovie := Response{}
	err := c.request(tmdbURL, mediaID, "POST", &removeMovie)
	if err != nil {
		return nil, err
	}
	return &removeMovie, nil
}

package tmdb

import (
	"fmt"
	"net/http"
)

// ListDetails type is a struct for details JSON response.
type ListDetails struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            string `json:"id"`
	Items         []struct {
		Adult            bool     `json:"adult,omitempty"` // Movie
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date,omitempty"` // TV
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name,omitempty"` // TV
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name,omitempty"`  // TV
		OriginalTitle    string   `json:"original_title,omitempty"` // Movie
		OriginCountry    []string `json:"origin_country,omitempty"` // TV
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		ReleaseDate      string   `json:"release_date,omitempty"` // Movie
		Title            string   `json:"title,omitempty"`        // Movie
		Video            bool     `json:"video,omitempty"`        // Movie
		VoteAverage      float32  `json:"vote_average"`
		VoteCount        int64    `json:"vote_count"`
	} `json:"items"`
	ItemCount  int64  `json:"item_count"`
	Iso639_1   string `json:"iso_639_1"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
}

// GetListDetails get the details of a list.
//
// https://developers.themoviedb.org/3/lists/get-list-details
func (c *Client) GetListDetails(
	id string,
	urlOptions map[string]string,
) (*ListDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%s?api_key=%s%s",
		baseURL,
		listURL,
		id,
		c.apiKey,
		options,
	)
	ListDetails := ListDetails{}
	if err := c.get(tmdbURL, &ListDetails); err != nil {
		return nil, err
	}
	return &ListDetails, nil
}

// ListItemStatus type is a struct for item status JSON response.
type ListItemStatus struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// GetListItemStatus check if a movie has already been added to the list.
//
// https://developers.themoviedb.org/3/lists/check-item-status
func (c *Client) GetListItemStatus(
	id string,
	urlOptions map[string]string,
) (*ListItemStatus, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%s/item_status?api_key=%s%s",
		baseURL,
		listURL,
		id,
		c.apiKey,
		options,
	)
	listItemStatus := ListItemStatus{}
	if err := c.get(tmdbURL, &listItemStatus); err != nil {
		return nil, err
	}
	return &listItemStatus, nil
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

// CreateList creates a list.
//
// https://developers.themoviedb.org/3/lists/create-list
func (c *Client) CreateList(
	list *ListCreate,
) (*ListResponse, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list?api_key=%s&session_id=%s",
		baseURL,
		c.apiKey,
		c.sessionID,
	)
	createList := ListResponse{}
	if err := c.request(
		tmdbURL,
		list,
		http.MethodPost,
		&createList,
	); err != nil {
		return nil, err
	}
	return &createList, nil
}

// ListMedia type is a struct for list media JSON request.
type ListMedia struct {
	MediaID int64 `json:"media_id"`
}

// AddMovie add a movie to a list.
//
// https://developers.themoviedb.org/3/lists/add-movie
func (c *Client) AddMovie(listID int64, mediaID *ListMedia) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d/add_item?api_key=%s&session_id=%s",
		baseURL,
		listID,
		c.apiKey,
		c.sessionID,
	)
	response := Response{}
	if err := c.request(
		tmdbURL,
		mediaID,
		http.MethodPost,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

// RemoveMovie remove a movie from a list.
//
// https://developers.themoviedb.org/3/lists/remove-movie
func (c *Client) RemoveMovie(listID int64, mediaID *ListMedia) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d/remove_item?api_key=%s&session_id=%s",
		baseURL,
		listID,
		c.apiKey,
		c.sessionID,
	)
	response := Response{}
	if err := c.request(
		tmdbURL,
		mediaID,
		http.MethodPost,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

// ClearList clear all of the items from a list.
//
// https://developers.themoviedb.org/3/lists/clear-list
func (c *Client) ClearList(listID int64, confirm bool) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d/clear?api_key=%s&session_id=%s&confirm=%t",
		baseURL,
		listID,
		c.apiKey,
		c.sessionID,
		confirm,
	)
	response := Response{}
	if err := c.request(
		tmdbURL,
		listID,
		http.MethodPost,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteList deletes a list.
//
// https://developers.themoviedb.org/3/lists/delete-list
func (c *Client) DeleteList(listID int64) (*Response, error) {
	tmdbURL := fmt.Sprintf(
		"%s/list/%d?api_key=%s&session_id=%s",
		baseURL,
		listID,
		c.apiKey,
		c.sessionID,
	)
	response := Response{}
	if err := c.request(
		tmdbURL,
		nil,
		http.MethodDelete,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

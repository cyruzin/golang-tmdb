package tmdb

import (
	"fmt"
	"net/http"
)

// ListDetails type is a struct for details JSON response.
type ListDetails struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int    `json:"favorite_count"`
	ID            int    `json:"id"`
	Items         []struct {
		Adult            bool     `json:"adult,omitempty"` // Movie
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date,omitempty"` // TV
		GenreIDs         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
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
		VoteMetrics
	} `json:"items"`
	ItemCount  int    `json:"item_count"`
	Iso639_1   string `json:"iso_639_1"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
}

// GetListDetails get the details of a list.
//
// https://developers.themoviedb.org/3/lists/get-list-details
func (c *Client) GetListDetails(id int, urlOptions map[string]string) (*ListDetails, error) {
	url := c.baseURL.JoinPath(listPath, fmt.Sprintf("%d", id))
	var ListDetails ListDetails
	if err := c.get(url, urlOptions, &ListDetails); err != nil {
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
func (c *Client) GetListItemStatus(id int, urlOptions map[string]string) (*ListItemStatus, error) {
	url := c.baseURL.JoinPath(listPath, fmt.Sprintf("%d", id), "item_status")
	var listItemStatus ListItemStatus
	if err := c.get(url, urlOptions, &listItemStatus); err != nil {
		return nil, err
	}
	return &listItemStatus, nil
}

// ListResponse type is a struct for list creation JSON response.
type ListResponse struct {
	*Response
	Success bool `json:"success"`
	ListID  int  `json:"list_id"`
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
func (c *Client) CreateList(list *ListCreate) (*ListResponse, error) {
	url := c.baseURL.JoinPath(listPath)
	urlOptions := map[string]string{}
	urlOptions["session_id"] = c.sessionID
	var createList ListResponse
	if err := c.request(url, http.MethodPost, urlOptions, list, &createList); err != nil {
		return nil, err
	}
	return &createList, nil
}

// ListMedia type is a struct for list media JSON request.
type ListMedia struct {
	MediaID int `json:"media_id"`
}

// AddMovie add a movie to a list.
//
// https://developers.themoviedb.org/3/lists/add-movie
func (c *Client) AddMovie(listID int, mediaID *ListMedia) (*Response, error) {
	url := c.baseURL.JoinPath(listPath, fmt.Sprintf("%d", listID), "add_item")
	urlOptions := map[string]string{}
	urlOptions["session_id"] = c.sessionID
	var response Response
	if err := c.request(url, http.MethodPost, urlOptions, mediaID, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RemoveMovie remove a movie from a list.
//
// https://developers.themoviedb.org/3/lists/remove-movie
func (c *Client) RemoveMovie(listID int, mediaID *ListMedia) (*Response, error) {
	url := c.baseURL.JoinPath(listPath, fmt.Sprintf("%d", listID), "remove_item")
	urlOptions := map[string]string{}
	urlOptions["session_id"] = c.sessionID
	var response Response
	if err := c.request(url, http.MethodPost, urlOptions, mediaID, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ClearList clear all of the items from a list.
//
// https://developers.themoviedb.org/3/lists/clear-list
func (c *Client) ClearList(
	listID int,
	confirm bool,
) (*Response, error) {
	url := c.baseURL.JoinPath(listPath, fmt.Sprintf("%d", listID), "clear")
	urlOptions := map[string]string{}
	urlOptions["session_id"] = c.sessionID
	var response Response
	if err := c.request(url, http.MethodPost, urlOptions, listID, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteList deletes a list.
//
// https://developers.themoviedb.org/3/lists/delete-list
func (c *Client) DeleteList(listID int) (*Response, error) {
	url := c.baseURL.JoinPath(listPath, fmt.Sprintf("%d", listID))
	urlOptions := map[string]string{}
	urlOptions["session_id"] = c.sessionID
	var response Response
	if err := c.request(url, http.MethodDelete, urlOptions, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

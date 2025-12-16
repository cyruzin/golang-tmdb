package tmdb

import (
	"fmt"
	"net/http"
)

// GetListDetails get the details of a list.
//
// https://developer.themoviedb.org/reference/list-details
func (c *Client) GetListDetails(
	id int64,
	urlOptions map[string]string,
) (*ListDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
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

// GetListItemStatus check if a movie has already been added to the list.
//
// https://developer.themoviedb.org/reference/list-check-item-status
func (c *Client) GetListItemStatus(
	id int64,
	urlOptions map[string]string,
) (*ListItemStatus, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/item_status?api_key=%s%s",
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

// CreateList creates a list.
//
// https://developer.themoviedb.org/reference/list-create
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

// AddMovie add a movie to a list.
//
// https://developer.themoviedb.org/reference/list-add-movie
func (c *Client) AddMovie(
	listID int,
	mediaID *ListMedia,
) (*Response, error) {
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
// https://developer.themoviedb.org/reference/list-remove-movie
func (c *Client) RemoveMovie(
	listID int,
	mediaID *ListMedia,
) (*Response, error) {
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
// https://developer.themoviedb.org/reference/list-clear
func (c *Client) ClearList(
	listID int,
	confirm bool,
) (*Response, error) {
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
// https://developer.themoviedb.org/reference/list-delete
func (c *Client) DeleteList(
	listID int,
) (*Response, error) {
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

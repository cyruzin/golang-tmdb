package tmdb

import (
	"fmt"
)

// Authentication type is a struct for authentication JSON response.
type Authentication struct {
	Success      bool   `json:"success"`
	ExpiresAt    string `json:"expires_at"`
	RequestToken string `json:"request_token"`
}

// Session type is a struct for session JSON response.
type Session struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// CreateRequestToken creates a temporary request token
// that can be used to validate a TMDb user login.
//
// Query String: api_key.
//
// https://developers.themoviedb.org/3/authentication/create-request-token
func (c *Client) CreateRequestToken() (*Authentication, error) {
	tmdbURL := fmt.Sprintf("%s%stoken/new?api_key=%s", baseURL, authenticationURL, c.APIKey)
	a := Authentication{}
	err := c.get(tmdbURL, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateSession creates a new session id.
//
// You can use this method to create a fully valid session ID
// once a user has validated the request token.
//
// Query String: api_key.
//
// https://developers.themoviedb.org/3/authentication/create-session
func (c *Client) CreateSession(requestToken string) (*Session, error) {
	tmdbURL := fmt.Sprintf("%s%ssession/new?api_key=%s", baseURL, authenticationURL, c.APIKey)
	rt := Authentication{RequestToken: requestToken}
	a := Session{}
	err := c.post(tmdbURL, &rt, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

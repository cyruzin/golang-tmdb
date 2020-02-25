package tmdb

import (
	"fmt"
)

// RequestToken type is a struct for request token JSON response.
type RequestToken struct {
	Success        bool   `json:"success"`
	ExpiresAt      string `json:"expires_at"`
	GuestSessionID string `json:"guest_session_id,omitempty"`
	RequestToken   string `json:"request_token,omitempty"`
}

// CreateGuestSession creates a temporary request token
// that can be used to validate a TMDb user login.
//
// https://developers.themoviedb.org/3/authentication/create-guest-session
func (c *Client) CreateGuestSession() (*RequestToken, error) {
	tmdbURL := fmt.Sprintf(
		"%s%sguest_session/new?api_key=%s",
		baseURL,
		authenticationURL,
		c.apiKey,
	)
	requestToken := RequestToken{}
	if err := c.get(tmdbURL, &requestToken); err != nil {
		return nil, err
	}
	return &requestToken, nil
}

// CreateRequestToken creates a temporary request token
// that can be used to validate a TMDb user login.
//
// https://developers.themoviedb.org/3/authentication/create-request-token
func (c *Client) CreateRequestToken() (*RequestToken, error) {
	tmdbURL := fmt.Sprintf(
		"%s%stoken/new?api_key=%s",
		baseURL,
		authenticationURL,
		c.apiKey,
	)
	requestToken := RequestToken{}
	if err := c.get(tmdbURL, &requestToken); err != nil {
		return nil, err
	}
	return &requestToken, nil
}

// AccessToken type is a struct for access token JSON request.
type AccessToken struct {
	AccessToken string `json:"access_token"`
}

// Session type is a struct for session JSON response.
type Session struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// SessionWithLogin type is a struct for session with login JSON response.
type SessionWithLogin struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	RequestToken string `json:"request_token"`
}

// CreateSession creates a new session id.
//
// You can use this method to create a fully valid session ID
// once a user has validated the request token.
//
// https://developers.themoviedb.org/3/authentication/create-session
// func (c *Client) CreateSession(rt string) (*Session, error) {
// 	tmdbURL := fmt.Sprintf("%s%ssession/new?api_key=%s", baseURL, authenticationURL, c.apiKey)
// 	requestToken := struct {
// 		RequestToken string `json:"request_token"`
// 	}{rt}
// 	j, err := json.Marshal(requestToken)
// 	if err != nil {
// 		return nil, err
// 	}
// 	a := Session{}
// 	err = c.post(tmdbURL, j, a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

// CreateSessionWithLogin creates a new session id using login.
//
// This method allows an application to validate a request token
// by entering a username and password. Not all applications have
// access to a web view so this can be used as a substitute.
// If you decide to use this method please use HTTPS.
//
// https://developers.themoviedb.org/3/authentication/validate-request-token
// func (c *Client) CreateSessionWithLogin(u, p, rt string) (*RequestToken, error) {
// 	tmdbURL := fmt.Sprintf("%s%stoken/validate_with_login?api_key=%s", baseURL, authenticationURL, c.apiKey)
// 	loginSession := SessionWithLogin{
// 		Username:     u,
// 		Password:     p,
// 		RequestToken: rt,
// 	}
// 	a := RequestToken{}
// 	err := c.post(tmdbURL, &loginSession, &a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

// CreateSessionFromV4 creates a new session id.
//
// Use this method to create a v3 session ID if you already have
// a valid v4 access token. The v4 token needs to be authenticated by the user.
// Your standard "read token" will not validate to create a session ID.
//
// https://developers.themoviedb.org/3/authentication/create-session-from-v4-access-token
// func (c *Client) CreateSessionFromV4(at string) (*Session, error) {
// 	tmdbURL := fmt.Sprintf("%s%ssession/convert/4?api_key=%s", baseURL, authenticationURL, c.apiKey)
// 	accessToken := AccessToken{AccessToken: at}
// 	a := Session{}
// 	err := c.post(tmdbURL, &accessToken, &a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

// TODO: Delete Session (DELETE Request)

package tmdb

// RequestToken type is a struct for request token JSON response.
type RequestToken struct {
	Success        bool   `json:"success"`
	ExpiresAt      string `json:"expires_at"`
	GuestSessionID string `json:"guest_session_id,omitempty"`
	RequestToken   string `json:"request_token,omitempty"`
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

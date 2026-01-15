// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tmdb is a Golang wrapper around the TMDb API
package tmdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	defaultBaseURL   *url.URL
	alternateBaseURL *url.URL
)

const (
	authenticationPath = "/authentication"
	moviePath          = "/movie"
	tvPath             = "/tv"
	tvSeasonPath       = "/season"
	personPath         = "/person"
	searchPath         = "/search"
	collectionPath     = "/collection"
	companyPath        = "/company"
	configurationPath  = "/configuration"
	creditPath         = "/credit"
	discoverPath       = "/discover"
	networkPath        = "/network"
	keywordPath        = "/keyword"
	genrePath          = "/genre"
	guestSessionPath   = "/guest_session"
	listPath           = "/list"
	accountPath        = "/account"
	watchProvidersPath = "/watch/providers"
)

func init() {
	var err error
	defaultBaseURL, err = url.Parse("https://api.themoviedb.org/3")
	if err != nil {
		panic(fmt.Errorf("defaultBaseURL failed to parse: %w", err))
	}
	alternateBaseURL, err = url.Parse("https://api.tmdb.org/3")
	if err != nil {
		panic(fmt.Errorf("alternateBaseURL failed to parse: %w", err))
	}
}

// customTransport wraps the base transport to modify requests.
type customTransport struct {
	base http.RoundTripper
	// client reference to read apiKey/bearerToken dynamically
	client *Client
}

// RoundTrip implements http.RoundTripper, adding authentication and headers to requests.
func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.client.apiKey == "" && t.client.bearerToken == "" {
		return nil, fmt.Errorf("apiKey and bearerToken is empty")
	}
	req.Header.Add("content-type", "application/json;charset=utf-8")
	if t.client.apiKey != "" {
		q := req.URL.Query()
		q.Add("api_key", t.client.apiKey)
		req.URL.RawQuery = q.Encode()
	}

	if t.client.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+t.client.bearerToken)
	}

	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}

	for {
		resp, err := base.RoundTrip(req)
		if err != nil {
			return nil, err
		}
		if t.client.autoRetry && shouldRetry(resp.StatusCode) {
			if err := resp.Body.Close(); err != nil {
				return nil, err
			}
			time.Sleep(retryDuration(resp))
			continue
		}
		return resp, nil
	}
}

// Client type is a struct to instantiate this pkg.
type Client struct {
	// TMDb apiKey to use the client.
	apiKey string
	// bearerToken will be used for v4 requests.
	bearerToken string
	// sessionId to use the client.
	sessionID string
	// Auto retry flag to indicates if the client
	// should retry the previous operation.
	autoRetry bool
	// http.Client for custom configuration.
	http http.Client
	// baseURL for API requests.
	baseURL *url.URL
}

// Response type is a struct for http responses.
type Response struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// ClientOption configures a Client.
type ClientOption func(*Client) error

// WithSessionID sets the session ID for authenticated requests.
func WithSessionID(sid string) ClientOption {
	return func(c *Client) error {
		if sid == "" {
			return fmt.Errorf("the session id is empty")
		}
		c.sessionID = sid
		return nil
	}
}

// WithHTTPClient sets a custom http.Client configuration.
func WithHTTPClient(httpClient http.Client) ClientOption {
	return func(c *Client) error {
		c.http = httpClient
		return nil
	}
}

// WithAutoRetry enables automatic retry on rate limit errors.
func WithAutoRetry() ClientOption {
	return func(c *Client) error {
		c.autoRetry = true
		return nil
	}
}

// WithTimeout sets the request timeout. Default is 10 seconds.
func WithTimeout(d time.Duration) ClientOption {
	return func(c *Client) error {
		c.http.Timeout = d
		return nil
	}
}

// WithAlternateBaseURL sets the alternate base URL.
func WithAlternateBaseURL() ClientOption {
	return func(c *Client) error {
		c.baseURL = alternateBaseURL
		return nil
	}
}

// WithCustomBaseURL sets a custom base URL.
func WithCustomBaseURL(customURL string) ClientOption {
	return func(c *Client) error {
		u, err := url.Parse(customURL)
		if err != nil {
			return err
		}

		c.baseURL = u
		return nil
	}
}

// WithAPIKey sets the API key for v3 authentication.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) error {
		if apiKey == "" {
			return fmt.Errorf("api key is empty")
		}
		c.apiKey = apiKey
		return nil
	}
}

// WithBearerToken sets the bearer token for v4 authentication.
func WithBearerToken(bearerToken string) ClientOption {
	return func(c *Client) error {
		if bearerToken == "" {
			return fmt.Errorf("bearer token is empty")
		}
		c.bearerToken = bearerToken
		return nil
	}
}

// New creates a new Client with the given options.
// At least one of WithAPIKey or WithBearerToken must be provided.
func New(options ...ClientOption) (*Client, error) {
	c := &Client{baseURL: defaultBaseURL}
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	if c.apiKey == "" && c.bearerToken == "" {
		return nil, fmt.Errorf("either API key or bearer token is required")
	}
	if c.http.Timeout == 0 {
		c.http.Timeout = time.Second * 10
	}
	c.http.Transport = &customTransport{base: c.http.Transport, client: c}
	return c, nil
}

// Auto retry default duration.
const defaultRetryDuration = time.Second * 5

// retryDuration calculates the retry duration time.
func retryDuration(resp *http.Response) time.Duration {
	retryTime := resp.Header.Get("Retry-After")
	if retryTime == "" {
		return defaultRetryDuration
	}
	seconds, err := strconv.ParseInt(retryTime, 10, 32)
	if err != nil {
		return defaultRetryDuration
	}
	return time.Duration(seconds) * time.Second
}

// shouldRetry determines whether the status code indicates that the
// previous operation should be retried at a later time.
func shouldRetry(status int) bool {
	return status == http.StatusAccepted || status == http.StatusTooManyRequests
}

func (c *Client) get(u *url.URL, opts map[string]string, dst any) error {
	if u == nil {
		return fmt.Errorf("url is nil")
	}
	if len(opts) > 0 {
		q := u.Query()
		for key, value := range opts {
			q.Add(key, url.QueryEscape(value))
		}
		u.RawQuery = q.Encode()
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("could not fetch the url: %s", err)
	}
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	if res.StatusCode != http.StatusOK {
		return c.decodeError(res)
	}
	if err = json.NewDecoder(res.Body).Decode(dst); err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}
	return nil
}

func (c *Client) request(u *url.URL, method string, opts map[string]string, body any, dst any) error {
	if u == nil {
		return fmt.Errorf("url is nil")
	}
	if len(opts) > 0 {
		q := u.Query()
		for key, value := range opts {
			q.Add(key, url.QueryEscape(value))
		}
		u.RawQuery = q.Encode()
	}
	bodyBytes := new(bytes.Buffer)
	if err := json.NewEncoder(bodyBytes).Encode(body); err != nil {
		return fmt.Errorf("failed to encode request body to JSON: %w", err)
	}
	req, err := http.NewRequestWithContext(
		context.Background(),
		method,
		u.String(),
		bytes.NewBuffer(bodyBytes.Bytes()),
	)
	if err != nil {
		return fmt.Errorf("could not fetch the url: %w", err)
	}
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Checking if the response is greater or equal
	// to 300 or less than 200.
	if res.StatusCode >= http.StatusMultipleChoices ||
		res.StatusCode < http.StatusOK ||
		res.StatusCode == http.StatusNoContent {
		return c.decodeError(res)
	}
	if err = json.NewDecoder(res.Body).Decode(dst); err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}
	return nil
}

// Error type represents an error returned by the TMDB API.
type Error struct {
	StatusMessage string `json:"status_message,omitempty"`
	Success       bool   `json:"success,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
}

// Error implements the error interface for API errors.
func (e Error) Error() string {
	return fmt.Sprintf(
		"code: %d | success: %t | message: %s",
		e.StatusCode,
		e.Success,
		e.StatusMessage,
	)
}

func (c *Client) decodeError(r *http.Response) error {
	resBody, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not read body response: %s", err)
	}
	if len(resBody) == 0 {
		return fmt.Errorf(
			"[%d]: empty body %s",
			r.StatusCode,
			http.StatusText(r.StatusCode),
		)
	}
	buf := bytes.NewBuffer(resBody)
	var clientError Error
	if err := json.NewDecoder(buf).Decode(&clientError); err != nil {
		return fmt.Errorf(
			"couldn't decode error: (%d) [%s]",
			len(resBody),
			resBody,
		)
	}
	return clientError
}

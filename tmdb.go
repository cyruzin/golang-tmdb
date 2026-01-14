// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tmdb is a Golang wrapper around the TMDb API
package tmdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	json "github.com/goccy/go-json"
)

// TMDb constants
const (
	defaultBaseURL    = "https://api.themoviedb.org/3"
	alternateBaseURL  = "https://api.tmdb.org/3"
	authenticationURL = "/authentication/"
	movieURL          = "/movie/"
	tvURL             = "/tv/"
	tvSeasonURL       = "/season/"
	tvEpisodeURL      = "/episode/"
	personURL         = "/person/"
	searchURL         = "/search/"
	collectionURL     = "/collection/"
	companyURL        = "/company/"
	configurationURL  = "/configuration/"
	creditURL         = "/credit/"
	discoverURL       = "/discover/"
	networkURL        = "/network/"
	keywordURL        = "/keyword/"
	genreURL          = "/genre/"
	guestSessionURL   = "/guest_session/"
	listURL           = "/list/"
	accountURL        = "/account/"
	watchProvidersURL = "/watch/providers/"
)

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
	baseURL string
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
			return errors.New("the session id is empty")
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

// WithAlternateBaseURL sets the alternate base URL.
func WithAlternateBaseURL() ClientOption {
	return func(c *Client) error {
		c.baseURL = alternateBaseURL
		return nil
	}
}

// WithCustomBaseURL sets a custom base URL.
func WithCustomBaseURL(url string) ClientOption {
	return func(c *Client) error {
		c.baseURL = url
		return nil
	}
}

// Init setups the Client with an apiKey.
func Init(apiKey string, options ...ClientOption) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("api key is empty")
	}
	c := &Client{apiKey: apiKey, baseURL: defaultBaseURL}
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// InitV4 setups the Client with an bearer token.
func InitV4(bearerToken string, options ...ClientOption) (*Client, error) {
	if bearerToken == "" {
		return nil, errors.New("bearer token is empty")
	}
	c := &Client{bearerToken: bearerToken, baseURL: defaultBaseURL}
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
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

func (c *Client) get(url string, data any) error {
	if url == "" {
		return errors.New("url field is empty")
	}
	if c.http.Timeout == 0 {
		c.http.Timeout = time.Second * 10
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.http.Timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("could not fetch the url: %s", err)
	}
	req.Header.Add("content-type", "application/json;charset=utf-8")
	if c.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	}
	for {
		res, err := c.http.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusTooManyRequests && c.autoRetry {
			time.Sleep(retryDuration(res))
			continue
		}
		if res.StatusCode == http.StatusNoContent {
			return nil
		}
		if res.StatusCode != http.StatusOK {
			return c.decodeError(res)
		}
		if err = json.NewDecoder(res.Body).Decode(data); err != nil {
			return fmt.Errorf("could not decode the data: %s", err)
		}
		break
	}
	return nil
}

func (c *Client) request(
	url string,
	body any,
	method string,
	data any,
) error {
	if url == "" {
		return errors.New("url field is empty")
	}
	if c.http.Timeout == 0 {
		c.http.Timeout = time.Second * 10
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.http.Timeout)
	defer cancel()
	bodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(bodyBytes).Encode(body)
	if err != nil {
		return fmt.Errorf("failed to encode request body to JSON: %s", err)
	}
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(bodyBytes.Bytes()),
	)
	if err != nil {
		return fmt.Errorf("could not fetch the url: %s", err)
	}
	req.Header.Add("content-type", "application/json;charset=utf-8")
	if c.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	}
	for {
		res, err := c.http.Do(req)
		if err != nil {
			return errors.New(err.Error())
		}
		defer res.Body.Close()
		if c.autoRetry && shouldRetry(res.StatusCode) {
			time.Sleep(retryDuration(res))
			continue
		}
		// Checking if the response is greater or equal
		// to 300 or less than 200.
		if res.StatusCode >= http.StatusMultipleChoices ||
			res.StatusCode < http.StatusOK ||
			res.StatusCode == http.StatusNoContent {
			return c.decodeError(res)
		}
		if err = json.NewDecoder(res.Body).Decode(data); err != nil {
			return fmt.Errorf("could not decode the data: %s", err)
		}
		break
	}
	return nil
}

func (c *Client) fmtOptions(
	urlOptions map[string]string,
) string {
	options := ""
	if len(urlOptions) > 0 {
		for key, value := range urlOptions {
			options += fmt.Sprintf(
				"&%s=%s",
				key,
				url.QueryEscape(value),
			)
		}
	}
	return options
}

// GetBaseURL gets the current base url.
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// Error type represents an error returned by the TMDB API.
type Error struct {
	StatusMessage string `json:"status_message,omitempty"`
	Success       bool   `json:"success,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
}

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

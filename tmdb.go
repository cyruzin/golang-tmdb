package tmdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client type is a struct for...
type Client struct {
	APIKey string
}

// Error type represents an error returned by the TMDB API.
type Error struct {
	StatusMessage string `json:"status_message,omitempty"`
	Success       bool   `json:"success,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
}

const (
	baseURL   = "https://api.themoviedb.org/3"
	imageURL  = "https://image.tmdb.org/t/p"
	movieURL  = "/movie/"
	tvURL     = "/tv/"
	peopleURL = "/people/"
)

func (c *Client) get(url string, data interface{}) error {

	if url == "" {
		return errors.New("url field is empty")
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.decodeError(res)
	}

	err = json.NewDecoder(res.Body).Decode(data)

	return err
}

func (c *Client) fmtOptions(o map[string]string) string {
	options := ""

	if len(o) > 0 {
		for k, v := range o {
			options += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	return options
}

func (e Error) Error() string {
	return e.StatusMessage
}

func (c *Client) decodeError(r *http.Response) error {
	resBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	if len(resBody) == 0 {
		return fmt.Errorf("[%d]: %s", r.StatusCode, http.StatusText(r.StatusCode))
	}

	buf := bytes.NewBuffer(resBody)

	var e struct {
		E Error `json:"error"`
	}

	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("couldn't decode error: (%d) [%s]", len(resBody), resBody)
	}

	if e.E.StatusMessage == "" {
		e.E.StatusMessage = fmt.Sprintf("[%d]: %s", r.StatusCode, http.StatusText(r.StatusCode))
	}

	return e.E
}

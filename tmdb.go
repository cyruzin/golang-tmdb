// Copyright (c) 2019 Cyro Dubeux. License MIT.

// Package tmdb is a wrapper for working with TMDb API.
package tmdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseURL           = "https://api.themoviedb.org/3"
	permissionURL     = "https://www.themoviedb.org/authenticate/"
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
)

// Client type is a struct to instantiate this pkg.
type Client struct {
	APIKey string
}

// Error type represents an error returned by the TMDB API.
type Error struct {
	StatusMessage string `json:"status_message,omitempty"`
	Success       bool   `json:"success,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
}

func (c *Client) get(url string, data interface{}) error {
	if url == "" {
		return errors.New("url field is empty")
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not fetch the url: %s", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return c.decodeError(res)
	}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}
	return nil
}

func (c *Client) post(url string, params []byte, data interface{}) error {
	if url == "" {
		return errors.New("url field is empty")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return errors.New(err.Error())
	}
	req.Header.Add("content-type", "application/json;charset=utf-8")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return c.decodeError(res)
	}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return fmt.Errorf("could not decode the data: %s", err)
	}
	return nil

}

func (c *Client) fmtOptions(o map[string]string) string {
	options := ""
	if len(o) > 0 {
		for k, v := range o {
			options += fmt.Sprintf("&%s=%s", k, url.QueryEscape(v))
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
		return fmt.Errorf("could not read body response: %s", err)
	}
	if len(resBody) == 0 {
		return fmt.Errorf("[%d]: empty body %s", r.StatusCode, http.StatusText(r.StatusCode))
	}
	buf := bytes.NewBuffer(resBody)
	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("couldn't decode error: (%d) [%s]", len(resBody), resBody)
	}
	return e
}

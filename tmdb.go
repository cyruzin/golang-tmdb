package tmdb

import (
	"io/ioutil"
	"net/http"
)

// Client type is a struct for...
type Client struct {
	APIKey string
}

const baseURL = "https://api.themoviedb.org/3"

func (c *Client) get(url string) (interface{}, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

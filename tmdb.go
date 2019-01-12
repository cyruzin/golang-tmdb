package tmdb

import (
	"encoding/json"
	"net/http"
)

// Client type is a struct for...
type Client struct {
	APIKey string
}

const baseURL = "https://api.themoviedb.org/3"

func (c *Client) get(url string, data interface{}) error {
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(data)

	if err != nil {
		return err
	}

	return nil
}

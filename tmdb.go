package tmdb

import (
	"encoding/json"
	"net/http"
)

// TMDB type is a struct for...
type TMDB struct {
	APIKey string
}

const baseURL = "https://api.themoviedb.org/3"

func (t *TMDB) get(url string, data interface{}) error {
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

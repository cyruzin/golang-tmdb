package tmdb

import "fmt"

// GetFindByID the find method makes it easy to search for objects in our
// database by an external id. For example, an IMDB ID.
//
// This method will search all objects (movies, TV shows and people)
// and return the results in a single response.
//
// https://developer.themoviedb.org/reference/find-by-id
func (c *Client) GetFindByID(
	id string,
	urlOptions map[string]string,
) (*FindByID, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s/find/%s?api_key=%s%s",
		baseURL, id, c.apiKey, options,
	)
	findByID := FindByID{}
	if err := c.get(tmdbURL, &findByID); err != nil {
		return nil, err
	}
	return &findByID, nil
}

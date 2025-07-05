package tmdb

import "fmt"

// FindByID type is a struct for find JSON response.
type FindByID struct {
	MovieResults     []MovieMedia     `json:"movie_results"`
	PersonResults    []PersonResults  `json:"person_results"`
	TvResults        []TVShowMedia    `json:"tv_results"`
	TvEpisodeResults []TVEpisodeMedia `json:"tv_episode_results"`
	TvSeasonResults  []TVSeasonMedia  `json:"tv_season_results"`
}

// GetFindByID the find method makes it easy to search for objects in our
// database by an external id. For example, an IMDB ID.
//
// This method will search all objects (movies, TV shows and people)
// and return the results in a single response.
//
// https://developers.themoviedb.org/3/find/find-by-id
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

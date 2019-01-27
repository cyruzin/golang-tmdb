package tmdb

import "fmt"

// TVEpisodeDetails type is a struct for details JSON response.
type TVEpisodeDetails struct {
	AirDate string `json:"air_date"`
	Crew    []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Department  string `json:"department"`
		Job         string `json:"job"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	EpisodeNumber int `json:"episode_number"`
	GuestStars    []struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		CreditID    string `json:"credit_id"`
		Character   string `json:"character"`
		Order       int    `json:"order"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ID             int64   `json:"id"`
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      int64   `json:"vote_count"`
}

// TVEpisodeChanges type is a struct for changes JSON response.
type TVEpisodeChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
			Value  string `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

// GetTVEpisodeDetails get the TV episode details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-details
func (c *Client) GetTVEpisodeDetails(id, s, e int, o map[string]string) (*TVEpisodeDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.APIKey, options,
	)
	t := TVEpisodeDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeChanges get the changes for a TV episode.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-changes
func (c *Client) GetTVEpisodeChanges(id int, o map[string]string) (*TVEpisodeChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%sepisode/%d/changes?api_key=%s%s",
		baseURL, tvURL, id, c.APIKey, options,
	)
	t := TVEpisodeChanges{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

package tmdb

import (
	"fmt"

	json "github.com/goccy/go-json"
)

// TVEpisodeGroupsDetails type is a struct for details JSON response.
type TVEpisodeGroupsDetails struct {
	Description  string `json:"description"`
	EpisodeCount int    `json:"episode_count"`
	GroupCount   int    `json:"group_count"`
	Groups       []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Order    int    `json:"order"`
		Episodes []struct {
			AirDate        string          `json:"air_date"`
			EpisodeNumber  int             `json:"episode_number"`
			ID             int64           `json:"id"`
			Name           string          `json:"name"`
			Overview       string          `json:"overview"`
			ProductionCode json.RawMessage `json:"production_code"`
			SeasonNumber   int             `json:"season_number"`
			ShowID         int64           `json:"show_id"`
			StillPath      string          `json:"still_path"`
			VoteAverage    float32         `json:"vote_average"`
			VoteCount      int64           `json:"vote_count"`
			Order          int             `json:"order"`
		} `json:"episodes"`
		Locked bool `json:"locked"`
	} `json:"groups"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Network struct {
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		Name          string `json:"name"`
		OriginCountry string `json:"origin_country"`
	} `json:"network"`
	Type int `json:"type"`
}

// GetTVEpisodeGroupsDetails the details of a TV episode group.
// Groups support 7 different types which are enumerated as the following:
//
// 1. Original air date
// 2. Absolute
// 3. DVD
// 4. Digital
// 5. Story arc
// 6. Production
// 7. TV
//
// https://developers.themoviedb.org/3/tv-episode-groups/get-tv-episode-group-details
func (c *Client) GetTVEpisodeGroupsDetails(
	id string,
	urlOptions map[string]string,
) (*TVEpisodeGroupsDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sepisode_group/%s?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvEpisodeGroupDetails := TVEpisodeGroupsDetails{}
	if err := c.get(tmdbURL, &tvEpisodeGroupDetails); err != nil {
		return nil, err
	}
	return &tvEpisodeGroupDetails, nil
}

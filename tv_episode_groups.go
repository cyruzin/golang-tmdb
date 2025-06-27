package tmdb

import (
	"fmt"
)

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
// https://developer.themoviedb.org/reference/tv-episode-group-details
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

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
	*TVEpisodeCreditsShort
	*TVEpisodeExternalIDsShort
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

// TVEpisodeCredits type is a struct for credits JSON response.
type TVEpisodeCredits struct {
	Cast []struct {
		Character   string `json:"character"`
		CreditID    string `json:"credit_id"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Department  string `json:"department"`
		Job         string `json:"job"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	GuestStars []struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		CreditID    string `json:"credit_id"`
		Character   string `json:"character"`
		Order       int    `json:"order"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	ID int64 `json:"id,omitempty"`
}

// TVEpisodeCreditsShort type is a short struct for credits JSON response.
type TVEpisodeCreditsShort struct {
	Credits *TVEpisodeCredits `json:"credits,omitempty"`
}

// TVEpisodeExternalIDs type is a struct for external ids JSON response.
type TVEpisodeExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	IMDbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
}

// TVEpisodeExternalIDsShort type is a short struct for external ids JSON response.
type TVEpisodeExternalIDsShort struct {
	ExternalIDs *TVEpisodeExternalIDs `json:"external_ids,omitempty"`
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

// GetTVEpisodeCredits get the credits (cast, crew and guest stars) for a TV episode.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-credits
func (c *Client) GetTVEpisodeCredits(id, s, e int) (*TVEpisodeCredits, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/credits?api_key=%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.APIKey,
	)
	t := TVEpisodeCredits{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeExternalIDs get the external ids for a TV episode.
// We currently support the following external sources.
//
// Media Databases: IMDb ID, TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-external-ids
func (c *Client) GetTVEpisodeExternalIDs(id, s, e int) (*TVEpisodeExternalIDs, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/external_ids?api_key=%s",
		baseURL, tvURL, id, tvSeasonURL, s, tvEpisodeURL, e, c.APIKey,
	)
	t := TVEpisodeExternalIDs{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// TODO: Account States

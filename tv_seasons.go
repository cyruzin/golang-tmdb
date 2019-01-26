package tmdb

import "fmt"

// TVSeasonsDetails is a struct for details JSON response.
type TVSeasonsDetails struct {
	IDString string `json:"_id"`
	AirDate  string `json:"air_date"`
	Episodes []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
		Crew           []struct {
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
	} `json:"episodes"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	ID           int64  `json:"id"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
	*TVSeasonsCreditsShort
}

// TVSeasonsChanges is a struct for changes JSON response.
type TVSeasonsChanges struct {
	Changes []struct {
		Items []struct {
			ID        string `json:"id"`
			Action    string `json:"action"`
			Time      string `json:"time"`
			Iso639_1  string `json:"iso_639_1"`
			Iso3166_1 string `json:"iso_3166_1"`
			Value     struct {
				EpisodeID     int64 `json:"episode_id"`
				EpisodeNumber int   `json:"episode_number"`
			} `json:"value"`
		} `json:"items"`
		Key string `json:"key"`
	} `json:"changes"`
}

// TVSeasonsCredits type is a struct for credits JSON response.
type TVSeasonsCredits struct {
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
		CreditID    string `json:"credit_id"`
		Department  string `json:"department"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Job         string `json:"job"`
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	ID int `json:"id"`
}

// TVSeasonsCreditsShort type is a short struct for credits JSON response.
type TVSeasonsCreditsShort struct {
	Credits *TVSeasonsCredits `json:"credits"`
}

// GetTVSeasonsDetails get the TV season details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
func (c *Client) GetTVSeasonsDetails(id, s int, o map[string]string) (*TVSeasonsDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d%s%d?api_key=%s%s", baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options)
	t := TVSeasonsDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsChanges get the changes for a TV season.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-changes
func (c *Client) GetTVSeasonsChanges(id int, o map[string]string) (*TVSeasonsChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%sseason/%d/changes?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVSeasonsChanges{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsCredits get the credits for TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-credits
func (c *Client) GetTVSeasonsCredits(id, s int, o map[string]string) (*TVSeasonsCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d%s%d/credits?api_key=%s%s", baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options)
	t := TVSeasonsCredits{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// TODO: Account States

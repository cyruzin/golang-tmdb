package tmdb

import (
	"encoding/json"
	"fmt"
)

// TVDetails type is a struct for details JSON response.
type TVDetails struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int64  `json:"id"`
		CreditID    string `json:"credit_id"`
		Name        string `json:"name"`
		Gender      int    `json:"gender"`
		ProfilePath string `json:"profile_path"`
	} `json:"created_by"`
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage         string   `json:"homepage"`
	ID               int64    `json:"id"`
	InProduction     bool     `json:"in_production"`
	Languages        []string `json:"languages"`
	LastAirDate      string   `json:"last_air_date"`
	LastEpisodeToAir struct {
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
	} `json:"last_episode_to_air"`
	Name             string `json:"name"`
	NextEpisodeToAir struct {
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
	} `json:"next_episode_to_air"`
	Networks []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"networks"`
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string   `json:"overview"`
	Popularity          float32  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		Name          string `json:"name"`
		ID            int64  `json:"id"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	Seasons []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int64  `json:"id"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	} `json:"seasons"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   int64   `json:"vote_count"`
	*TVAlternativeTitlesShort
	*TVChangesShort
	*TVContentRatingsShort
	*TVCreditsShort
	*TVEpisodeGroupsShort
}

// TVAccountStates type is a struct for account states JSON response.
type TVAccountStates struct {
	ID        int64           `json:"id"`
	Favorite  bool            `json:"favorite"`
	Rated     json.RawMessage `json:"rated"`
	Watchlist bool            `json:"watchlist"`
}

// TVAlternativeTitle type is a struct for alternative title JSON response.
type TVAlternativeTitle struct {
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string `json:"title"`
		Type      string `json:"type"`
	} `json:"results"`
}

// TVAlternativeTitles type is a struct for alternative titles JSON response.
type TVAlternativeTitles struct {
	ID int `json:"id"`
	*TVAlternativeTitle
}

// TVAlternativeTitlesShort type is a short struct for alternative titles JSON response.
type TVAlternativeTitlesShort struct {
	AlternativeTitles *TVAlternativeTitle `json:"alternative_titles,omitempty"`
}

// TVChanges type is a struct for changes JSON response.
type TVChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID     string `json:"id"`
			Action string `json:"action"`
			Time   string `json:"time"`
			Value  struct {
				SeasonID     int64 `json:"season_id"`
				SeasonNumber int   `json:"season_number"`
			} `json:"value"`
		} `json:"items"`
	} `json:"changes"`
}

// TVChangesShort type is a short struct for changes JSON response.
type TVChangesShort struct {
	Changes *TVChanges `json:"changes,omitempty"`
}

// TVContentRating type a struct for content rating JSON response.
type TVContentRating struct {
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Rating    string `json:"rating"`
	} `json:"results"`
}

// TVContentRatings type is a struct for content ratings JSON response.
type TVContentRatings struct {
	*TVContentRating
	ID int64 `json:"id"`
}

// TVContentRatingsShort type is a short struct for content ratings JSON response.
type TVContentRatingsShort struct {
	ContentRatings *TVContentRating `json:"content_ratings,omitempty"`
}

// TVCreditsCast type is a struct for cast JSON response.
type TVCreditsCast struct {
	Character   string `json:"character"`
	CreditID    string `json:"credit_id"`
	Gender      int    `json:"gender"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profile_path"`
}

// TVCreditsCrew type is a struct for crew JSON response.
type TVCreditsCrew struct {
	CreditID    string `json:"credit_id"`
	Department  string `json:"department"`
	Gender      int    `json:"gender"`
	ID          int64  `json:"id"`
	Job         string `json:"job"`
	Name        string `json:"name"`
	ProfilePath string `json:"profile_path"`
}

// TVCredits type is a struct for credits JSON response.
type TVCredits struct {
	ID   int64 `json:"id"`
	Cast []struct {
		*TVCreditsCast
	} `json:"cast"`
	Crew []struct {
		*TVCreditsCrew
	} `json:"crew"`
}

// TVCreditsShort type is a short struct for credits JSON response.
type TVCreditsShort struct {
	Credits struct {
		Cast []*TVCreditsCast `json:"cast,omitempty"`
		Crew []*TVCreditsCrew `json:"crew,omitempty"`
	} `json:"credits,omitempty"`
}

// TVEpisodeGroup type is a struct for episode group JSON response.
type TVEpisodeGroup struct {
	Results []struct {
		Description  string `json:"description"`
		EpisodeCount int    `json:"episode_count"`
		GroupCount   int    `json:"group_count"`
		ID           string `json:"id"`
		Name         string `json:"name"`
		Network      []struct {
			ID            int64  `json:"id"`
			LogoPath      string `json:"logo_path"`
			Name          string `json:"name"`
			OriginCountry string `json:"origin_country"`
		} `json:"network"`
		Type int `json:"type"`
	} `json:"results"`
}

// TVEpisodeGroups type is a struct for episode groups JSON response.
type TVEpisodeGroups struct {
	*TVEpisodeGroup
	ID int64 `json:"id"`
}

// TVEpisodeGroupsShort type is a short struct for episode groups JSON response.
type TVEpisodeGroupsShort struct {
	EpisodeGroups *TVEpisodeGroup `json:"episode_groups,omitempty"`
}

// GetTVDetails get the primary TV show details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv/get-tv-details
func (c *Client) GetTVDetails(id int, o map[string]string) (*TVDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVAccountStates grab the following account states for a session:
//
// TV show rating.
//
// If it belongs to your watchlist.
//
// If it belongs to your favourite list.
//
// https://developers.themoviedb.org/3/tv/get-tv-account-states
//
func (c *Client) GetTVAccountStates(id int, o map[string]string) (*TVAccountStates, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/account_states?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVAccountStates{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVAlternativeTitles get all of the alternative titles for a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-alternative-titles
func (c *Client) GetTVAlternativeTitles(id int, o map[string]string) (*TVAlternativeTitles, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/alternative_titles?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVAlternativeTitles{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVChanges get the changes for a TV show.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// TV show changes are different than movie changes in that there
// are some edits on seasons and episodes that will create a change
// entry at the show level. These can be found under the season
// and episode keys. These keys will contain a series_id and episode_id.
// You can use the and methods to look these up individually.
//
// https://developers.themoviedb.org/3/tv/get-tv-changes
func (c *Client) GetTVChanges(id int, o map[string]string) (*TVChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/changes?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVChanges{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVContentRatings get the list of content ratings (certifications) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-content-ratings
func (c *Client) GetTVContentRatings(id int, o map[string]string) (*TVContentRatings, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/content_ratings?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVContentRatings{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVCredits get the credits (cast and crew) that have been added to a TV show.
//
// https://developers.themoviedb.org/3/tv/get-tv-credits
func (c *Client) GetTVCredits(id int, o map[string]string) (*TVCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/credits?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVCredits{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVEpisodeGroups get all of the episode groups that have been created for a TV show.
//
// With a group ID you can call the get TV episode group details method.
//
// https://developers.themoviedb.org/3/tv/get-tv-episode-groups
func (c *Client) GetTVEpisodeGroups(id int, o map[string]string) (*TVEpisodeGroups, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%s%d/episode_groups?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVEpisodeGroups{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

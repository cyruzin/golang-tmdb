package tmdb

import json "github.com/goccy/go-json"

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
			Order          int             `json:"order"`
			VoteMetrics
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

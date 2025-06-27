package tmdb

// TVEpisodeGroupsDetails type is a struct for details JSON response.
type TVEpisodeGroupsDetails struct {
	Description  string      `json:"description"`
	EpisodeCount int         `json:"episode_count"`
	GroupCount   int         `json:"group_count"`
	Groups       []Group     `json:"groups"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Network      CompanyInfo `json:"network"`
	Type         GroupType   `json:"type"`
}

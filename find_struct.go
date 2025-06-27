package tmdb

// FindByID type is a struct for find JSON response.
type FindByID struct {
	MovieResults     []MovieMedia     `json:"movie_results"`
	PersonResults    []PersonMedia    `json:"person_results"`
	TvResults        []TVShowMedia    `json:"tv_results"`
	TvEpisodeResults []TVEpisodeMedia `json:"tv_episode_results"`
	TvSeasonResults  []TVSeasonMedia  `json:"tv_season_results"`
}

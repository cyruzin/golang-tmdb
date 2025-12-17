package tmdb

// FindByID type is a struct for find JSON response.
type FindByID struct {
	MovieResults []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		ID               int64   `json:"id"`
		Title            string  `json:"title"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		MediaType        string  `json:"media_type"`
		GenreIDs         []int64 `json:"genre_ids"`
		Popularity       float32 `json:"popularity"`
		ReleaseDate      string  `json:"release_date"`
		Video            bool    `json:"video"`
		VoteMetrics
	} `json:"movie_results"`
	PersonResults []struct {
		Adult              bool    `json:"adult"`
		ID                 int64   `json:"id"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		MediaType          string  `json:"media_type"`
		Popularity         float32 `json:"popularity"`
		Gender             int     `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		ProfilePath        string  `json:"profile_path"`
		KnownFor           []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			ID               int64   `json:"id"`
			Title            string  `json:"title"`
			OriginalTitle    string  `json:"original_title"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			GenreIDs         []int   `json:"genre_ids"`
			Popularity       float32 `json:"popularity"`
			ReleaseDate      string  `json:"release_date"` // Movie
			Video            bool    `json:"video"`        // Movie
			VoteMetrics

			// FirstAirDate  string   `json:"first_air_date"` // TV
			// Name          string   `json:"name"`           // TV
			// OriginalName  string   `json:"original_name"`  // TV
			// OriginCountry []string `json:"origin_country"` // TV
		} `json:"known_for"`
	} `json:"person_results"`
	TvResults []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		ID               int64   `json:"id"`
		Name             string  `json:"name"`
		OriginalName     string  `json:"original_name"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		GenreIDs         []int64 `json:"genre_ids"`
		Popularity       float32 `json:"popularity"`
		FirstAirDate     string  `json:"first_air_date"`
		VoteMetrics
		OriginCountry []string `json:"origin_country"`
	} `json:"tv_results"`
	TvEpisodeResults []struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Overview  string `json:"overview"`
		MediaType string `json:"media_type"`
		VoteMetrics
		AirDate        string `json:"air_date"`
		EpisodeNumber  int    `json:"episode_number"`
		EpisodeType    string `json:"episode_type"`
		ProductionCode string `json:"production_code"`
		Runtime        int    `json:"runtime"`
		SeasonNumber   int    `json:"season_number"`
		ShowID         int64  `json:"show_id"`
		StillPath      string `json:"still_path"`
	} `json:"tv_episode_results"`
	TvSeasonResults []struct {
		ID           int64   `json:"id"`
		Name         string  `json:"name"`
		Overview     string  `json:"overview"`
		PosterPath   string  `json:"poster_path"`
		MediaType    string  `json:"media_type"`
		VoteAverage  float32 `json:"vote_average"`
		AirDate      string  `json:"air_date"`
		SeasonNumber int     `json:"season_number"`
		ShowID       int64   `json:"show_id"`
		EpisodeCount int     `json:"episode_count"`
	} `json:"tv_season_results"`
}

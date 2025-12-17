package tmdb

// GuestSessionRatedMovies type is a struct for rated movies JSON response.
type GuestSessionRatedMovies struct {
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		PosterPath       string  `json:"poster_path"`
		Popularity       float32 `json:"popularity"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		Rating           float32 `json:"rating"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

// GuestSessionRatedTVShows type is a struct for rated tv shows JSON response.
type GuestSessionRatedTVShows struct {
	Results []struct {
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		PosterPath       string   `json:"poster_path"`
		Popularity       float32  `json:"popularity"`
		Name             string   `json:"name"`
		Rating           float32  `json:"rating"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

// GuestSessionRatedTVEpisodes type is a struct for rated tv episodes JSON response.
type GuestSessionRatedTVEpisodes struct {
	Results []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		Rating         float32 `json:"rating"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

package tmdb

// CreditsDetails type is a struct for credits JSON response.
type CreditsDetails struct {
	CreditType string `json:"credit_type"`
	Department string `json:"department"`
	Job        string `json:"job"`
	Media      struct {
		Adult            bool    `json:"adult,omitempty"` // Movie
		BackdropPath     string  `json:"backdrop_path"`
		ID               int64   `json:"id"`
		Title            string  `json:"title,omitempty"`          // Movie
		OriginalTitle    string  `json:"original_title,omitempty"` // Movie
		Name             string  `json:"name,omitempty"`           // TV
		OriginalName     string  `json:"original_name,omitempty"`  // TV
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		GenreIDs         []int64 `json:"genre_ids"`
		Popularity       float32 `json:"popularity"`
		ReleaseDate      string  `json:"release_date,omitempty"`   // Movie
		Video            bool    `json:"video,omitempty"`          // Movie
		FirstAirDate     string  `json:"first_air_date,omitempty"` // TV
		VoteMetrics
		OriginCountry []string `json:"origin_country,omitempty"` // TV
		Character     string   `json:"character"`
		Episodes      []struct {
			ID        int64  `json:"id"`
			Name      string `json:"name"` // TV
			Overview  string `json:"overview"`
			MediaType string `json:"media_type"`
			VoteMetrics
			AirDate        string `json:"air_date"`
			EpisodeNumber  int64  `json:"episode_number"`
			EpisodeType    string `json:"episode_type"`
			ProductionCode string `json:"production_code"`
			Runtime        int    `json:"runtime"`
			SeasonNumber   int    `json:"season_number"`
			ShowID         int    `json:"show_id"`
			StillPath      string `json:"still_path"`
		} `json:"episodes,omitempty"` // TV
		Seasons []Season `json:"seasons,omitempty"` // TV
	} `json:"media"`
	MediaType string `json:"media_type"`
	ID        string `json:"id"`
	Person    struct {
		Adult              bool    `json:"adult"`
		ID                 int64   `json:"id"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name,omitempty"`
		MediaType          string  `json:"media_type"`
		Popularity         float32 `json:"popularity"`
		Gender             int     `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		ProfilePath        string  `json:"profile_path"`
		// KnownFor           []struct {
		// 	Adult        bool   `json:"adult,omitempty"`
		// 	BackdropPath string `json:"backdrop_path"`
		// 	// GenreIDs         []int64  `json:"genre_ids"` // FIXME: -> []float32
		// 	// ID               int64    `json:"id"` // FIXME: -> float32
		// 	OriginalLanguage string   `json:"original_language"`
		// 	OriginalTitle    string   `json:"original_title,omitempty"`
		// 	Overview         string   `json:"overview"`
		// 	PosterPath       string   `json:"poster_path"`
		// 	ReleaseDate      string   `json:"release_date,omitempty"`
		// 	Title            string   `json:"title,omitempty"`
		// 	Video            bool     `json:"video,omitempty"`
		// 	Popularity       float32  `json:"popularity"`
		// 	MediaType        string   `json:"media_type"`
		// 	OriginalName     string   `json:"original_name,omitempty"`
		// 	Name             string   `json:"name,omitempty"`
		// 	FirstAirDate     string   `json:"first_air_date,omitempty"`
		// 	OriginCountry    []string `json:"origin_country,omitempty"`
		// 	VoteMetrics
		// } `json:"known_for"`
	} `json:"person"`
}

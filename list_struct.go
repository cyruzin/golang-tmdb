package tmdb

// ListDetails type is a struct for details JSON response.
type ListDetails struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	Items         []struct {
		Adult            bool     `json:"adult,omitempty"` // Movie
		BackdropPath     string   `json:"backdrop_path"`
		FirstAirDate     string   `json:"first_air_date,omitempty"` // TV
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name,omitempty"` // TV
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name,omitempty"`  // TV
		OriginalTitle    string   `json:"original_title,omitempty"` // Movie
		OriginCountry    []string `json:"origin_country,omitempty"` // TV
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		ReleaseDate      string   `json:"release_date,omitempty"` // Movie
		Title            string   `json:"title,omitempty"`        // Movie
		Video            bool     `json:"video,omitempty"`        // Movie
		VoteMetrics
	} `json:"items"`
	ItemCount  int64  `json:"item_count"`
	Iso639_1   string `json:"iso_639_1"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
}

// ListItemStatus type is a struct for item status JSON response.
type ListItemStatus struct {
	ID          string `json:"id"`
	ItemPresent bool   `json:"item_present"`
}

// ListResponse type is a struct for list creation JSON response.
type ListResponse struct {
	*Response
	Success bool  `json:"success"`
	ListID  int64 `json:"list_id"`
}

// ListCreate type is a struct for list creation JSON request.
type ListCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
}

// ListMedia type is a struct for list media JSON request.
type ListMedia struct {
	MediaID int64 `json:"media_id"`
}

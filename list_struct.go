package tmdb

// ListDetails type is a struct for details JSON response.
type ListDetails struct {
	CreatedBy     string `json:"created_by"`
	Description   string `json:"description"`
	FavoriteCount int64  `json:"favorite_count"`
	ID            int64  `json:"id"`
	Iso639_1      string `json:"iso_639_1"`
	ItemCount     int64  `json:"item_count"`
	Items         []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		ID               int64   `json:"id"`
		Name             string  `json:"name"`
		OriginalName     string  `json:"original_name"`
		Title            string  `json:"title"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		GenreIDs         []int   `json:"genre_ids"`
		Popularity       float32 `json:"popularity"`
		FirstAirDate     string  `json:"first_air_date"`
		ReleaseDate      string  `json:"release_date"`
		Video            bool    `json:"video"`
		VoteMetrics
		OriginCountry []string `json:"origin_country"`
	} `json:"items"`
	Name       string `json:"name"`
	PosterPath string `json:"poster_path"`
	PaginatedResultsMeta
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

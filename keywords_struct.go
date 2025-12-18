package tmdb

// KeywordDetails type is a struct for keyword JSON response.
type KeywordDetails struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// KeywordMovies type is a struct for movies that belong to a keyword JSON response.
type KeywordMovies struct {
	ID      int64 `json:"id"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		Popularity       float32 `json:"popularity"`
		VoteMetrics
	} `json:"results"`
	PaginatedResultsMeta
}

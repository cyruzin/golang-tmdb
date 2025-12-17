package tmdb

// CollectionDetails type is a struct for details JSON response.
type CollectionDetails struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
	Parts        []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		Popularity       float32 `json:"popularity"`
		VoteMetrics
	} `json:"parts"`
}

// CollectionImage type is a struct for a single image.
type CollectionImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
}

// CollectionImages type is a struct for images JSON response.
type CollectionImages struct {
	ID        int64             `json:"id"`
	Backdrops []CollectionImage `json:"backdrops"`
	Posters   []CollectionImage `json:"posters"`
}

// CollectionTranslations type is a struct for translations JSON response.
type CollectionTranslations struct {
	ID           int64         `json:"id"`
	Translations []Translation `json:"translations"`
}

package tmdb

// CollectionDetails type is a struct for details JSON response.
type CollectionDetails struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	OriginalLanguage string `json:"original_language"`
	OriginalName     string `json:"original_name"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
	BackdropPath     string `json:"backdrop_path"`
	Parts            []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		ID               int64   `json:"id"`
		Title            string  `json:"title"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		GenreIDs         []int64 `json:"genre_ids"`
		Popularity       float32 `json:"popularity"`
		ReleaseDate      string  `json:"release_date"`
		Video            bool    `json:"video"`
		VoteMetrics
	} `json:"parts"`
}

// CollectionImage type is a struct for a single image.
type CollectionImage struct {
	ImageBase
	Iso3166_1 string `json:"iso_3166_1"`
	Iso639_1  string `json:"iso_639_1"`
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

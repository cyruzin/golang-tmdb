package tmdb

// CollectionDetails type is a struct for details JSON response.
type CollectionDetails struct {
	ID               int64        `json:"id"`
	Name             string       `json:"name"`
	OriginalLanguage string       `json:"original_language"`
	OriginalName     string       `json:"original_name"`
	Overview         string       `json:"overview"`
	PosterPath       string       `json:"poster_path"`
	BackdropPath     string       `json:"backdrop_path"`
	Parts            []MovieMedia `json:"parts"`
}

// CollectionImages type is a struct for images JSON response.
type CollectionImages struct {
	ID        int64      `json:"id"`
	Backdrops []ImageIso `json:"backdrops"`
	Posters   []ImageIso `json:"posters"`
}

// CollectionTranslations type is a struct for translations JSON response.
type CollectionTranslations struct {
	ID           int64                   `json:"id"`
	Translations []CollectionTranslation `json:"translations"`
}

type CollectionTranslationData struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
}

type CollectionTranslation struct {
	TranslationBase
	Data CollectionTranslationData `json:"data"`
}

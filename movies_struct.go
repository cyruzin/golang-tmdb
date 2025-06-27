package tmdb

// MovieDetails type is a struct for movie details JSON response.
type MovieDetails struct {
	VideoDetails
	MovieMeta
	OriginCountry       []string             `json:"origin_country"`
	BelongsToCollection *BelongsToCollection `json:"belongs_to_collection"`
	Budget              int64                `json:"budget"`
	IMDbID              string               `json:"imdb_id"`
	Revenue             int64                `json:"revenue"`
	Runtime             int                  `json:"runtime"`
	*MovieAlternativeTitlesAppend
	*MovieChangesAppend
	*MovieCreditsAppend
	*MovieExternalIDsAppend
	*MovieImagesAppend
	*MovieKeywordsAppend
	*MovieReleaseDatesAppend
	*MovieVideosAppend
	*MovieTranslationsAppend
	*MovieRecommendationsAppend
	*MovieSimilarAppend
	*MovieReviewsAppend
	*MovieListsAppend
	*MovieWatchProvidersAppend
}

// MovieAlternativeTitlesAppend type is a struct for alternative
// titles in append to response.
type MovieAlternativeTitlesAppend struct {
	AlternativeTitles MovieAlternativeTitles `json:"alternative_titles"`
}

// MovieChangesAppend type is a struct for changes in append to response.
type MovieChangesAppend struct {
	Changes Changes `json:"changes"`
}

// MovieCreditsAppend type is a struct for credits in append to response.
type MovieCreditsAppend struct {
	Credits MovieCredits `json:"credits"`
}

// MovieExternalIDsAppend type is a struct for external ids in append to response.
type MovieExternalIDsAppend struct {
	ExternalIDs MovieExternalIDs `json:"external_ids"`
}

// MovieImagesAppend type is a struct for images in append to response.
type MovieImagesAppend struct {
	Images MovieImages `json:"images"`
}

// MovieReleaseDatesAppend type is a struct for release dates in append to response.
type MovieReleaseDatesAppend struct {
	ReleaseDates MovieReleaseDateResults `json:"release_dates"`
}

// MovieVideosAppend type is a struct for videos in append to response.
type MovieVideosAppend struct {
	Videos VideoResults `json:"videos"`
}

// MovieTranslationsAppend type is a struct for translations in append to response.
type MovieTranslationsAppend struct {
	Translations MovieTranslations `json:"translations"`
}

// MovieRecommendationsAppend type is a struct for
// recommendations in append to response.
type MovieRecommendationsAppend struct {
	Recommendations PaginatedMovieResults `json:"recommendations"`
}

// MovieSimilarAppend type is a struct for similar movies in append to response.
type MovieSimilarAppend struct {
	Similar PaginatedMovieResults `json:"similar"`
}

// MovieReviewsAppend type is a struct for reviews in append to response.
type MovieReviewsAppend struct {
	Reviews PaginatedReviewsResults `json:"reviews"`
}

// MovieListsAppend type is a struct for lists in append to response.
type MovieListsAppend struct {
	Lists PaginatedMovieListResults `json:"lists"`
}

// MovieKeywordsAppend type is a struct for keywords in append to response.
type MovieKeywordsAppend struct {
	Keywords KeywordsList `json:"keywords"`
}

// MovieWatchProvidersAppend type is a struct for
// watch/providers in append to response.
type MovieWatchProvidersAppend struct {
	WatchProviders IDWatchProviderResults `json:"watch/providers"`
}

// MovieAccountStates type is a struct for account states JSON response.
type MovieAccountStates struct {
	AccountStates
}

// MovieAlternativeTitles type is a struct for alternative titles JSON response.
type MovieAlternativeTitles struct {
	ID     int                `json:"id"`
	Titles []AlternativeTitle `json:"titles"`
}

// MovieExternalIDs type is a struct for external ids JSON response.
type MovieExternalIDs struct {
	ID          int64   `json:"id"`
	IMDbID      *string `json:"imdb_id"`
	WikidataID  *string `json:"wikidata_id"`
	FacebookID  *string `json:"facebook_id"`
	InstagramID *string `json:"instagram_id"`
	TwitterID   *string `json:"twitter_id"`
}

// MovieImages type is a struct for images JSON response.
type MovieImages struct {
	ID        int64      `json:"id"`
	Backdrops []ImageIso `json:"backdrops"`
	Logos     []ImageIso `json:"logos"`
	Posters   []ImageIso `json:"posters"`
}

// MovieTranslations type is a struct for translations JSON response.
type MovieTranslations struct {
	ID           int64              `json:"id"`
	Translations []MovieTranslation `json:"translations"`
}

type MovieTranslation struct {
	TranslationBase
	Data MovieTranslationData `json:"data"`
}

type MovieTranslationData struct {
	Homepage string `json:"homepage"`
	Overview string `json:"overview"`
	Runtime  int    `json:"runtime"`
	Tagline  string `json:"tagline"`
	Title    string `json:"title"`
}

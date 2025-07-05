package tmdb

// MovieDetails type is a struct for movie details JSON response.
type MovieDetails struct {
	VideoDetails
	MovieMeta
	OriginCountry       []string            `json:"origin_country"`
	BelongsToCollection BelongsToCollection `json:"belongs_to_collection"`
	Budget              int64               `json:"budget"`
	IMDbID              string              `json:"imdb_id"`
	Revenue             int64               `json:"revenue"`
	Runtime             int                 `json:"runtime"`
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
	AlternativeTitles *MovieAlternativeTitles `json:"alternative_titles"`
}

// MovieChangesAppend type is a struct for changes in append to response.
type MovieChangesAppend struct {
	Changes *MovieChanges `json:"changes"`
}

// MovieCreditsAppend type is a struct for credits in append to response.
type MovieCreditsAppend struct {
	Credits *MovieCredits `json:"credits"`
}

// MovieExternalIDsAppend type is a struct for external ids in append to response.
type MovieExternalIDsAppend struct {
	*MovieExternalIDs `json:"external_ids"`
}

// MovieImagesAppend type is a struct for images in append to response.
type MovieImagesAppend struct {
	Images *MovieImages `json:"images"`
}

// MovieReleaseDatesAppend type is a struct for release dates in append to response.
type MovieReleaseDatesAppend struct {
	ReleaseDates *MovieReleaseDates `json:"release_dates"`
}

// MovieVideosAppend type is a struct for videos in append to response.
type MovieVideosAppend struct {
	Videos *VideoResults `json:"videos"`
}

// MovieTranslationsAppend type is a struct for translations in append to response.
type MovieTranslationsAppend struct {
	Translations *MovieTranslations `json:"translations"`
}

// MovieRecommendationsAppend type is a struct for
// recommendations in append to response.
type MovieRecommendationsAppend struct {
	Recommendations *MovieRecommendations `json:"recommendations"`
}

// MovieSimilarAppend type is a struct for similar movies in append to response.
type MovieSimilarAppend struct {
	Similar *MovieSimilar `json:"similar"`
}

// MovieReviewsAppend type is a struct for reviews in append to response.
type MovieReviewsAppend struct {
	Reviews *MovieReviews `json:"reviews,"`
}

// MovieListsAppend type is a struct for lists in append to response.
type MovieListsAppend struct {
	Lists *MovieLists `json:"lists,omitempty"`
}

// MovieKeywordsAppend type is a struct for keywords in append to response.
type MovieKeywordsAppend struct {
	Keywords *MovieKeywords `json:"keywords"`
}

// MovieWatchProvidersAppend type is a struct for
// watch/providers in append to response.
type MovieWatchProvidersAppend struct {
	WatchProviders *WatchProviderResults `json:"watch/providers"`
}

type RatedValue struct {
	Value float32 `json:"value"`
}

// MovieAccountStates type is a struct for account states JSON response.
type MovieAccountStates struct {
	ID        int64      `json:"id"`
	Favorite  bool       `json:"favorite"`
	Rated     RatedValue `json:"rated"`
	Watchlist bool       `json:"watchlist"`
}

// MovieAlternativeTitles type is a struct for alternative titles JSON response.
type MovieAlternativeTitles struct {
	ID     int                `json:"id"`
	Titles []AlternativeTitle `json:"titles"`
}

// MovieChanges type is a struct for changes JSON response.
type MovieChanges struct {
	Changes []Change `json:"changes"`
}

// MovieExternalIDs type is a struct for external ids JSON response.
type MovieExternalIDs struct {
	IMDbID      string `json:"imdb_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
	WikiDataID  string `json:"wikidata_id"`
	ID          int64  `json:"id,omitempty"`
}

// MovieImage type is a struct for a single image.
type MovieImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
}

// MovieImages type is a struct for images JSON response.
type MovieImages struct {
	ID        int64        `json:"id"`
	Backdrops []MovieImage `json:"backdrops"`
	Logos     []MovieImage `json:"logos"`
	Posters   []MovieImage `json:"posters"`
}

// MovieKeywords type is a struct for keywords JSON response.
type MovieKeywords struct {
	ID       int64            `json:"id"`
	Keywords []KeywordDetails `json:"keywords"`
}

// MovieReleaseDates type is a struct for release dates JSON response.
type MovieReleaseDates struct {
	ID int64 `json:"id"`
	*MovieReleaseDatesResults
}

// MovieTranslations type is a struct for translations JSON response.
type MovieTranslations struct {
	ID           int64         `json:"id"`
	Translations []Translation `json:"translations"`
}

// MovieRecommendations type is a struct for recommendations JSON response.
type MovieRecommendations struct {
	*MovieRecommendationsResults
	PaginatedResultsMeta
}

// MovieSimilar type is a struct for similar movies JSON response.
type MovieSimilar struct {
	*MovieRecommendations
}

// MovieReviews type is a struct for reviews JSON response.
type MovieReviews struct {
	ID int64 `json:"id"`
	*MovieReviewsResults
	PaginatedResultsMeta
}

// MovieLists type is a struct for lists JSON response.
type MovieLists struct {
	ID int64 `json:"id"`
	*MovieListsResults
	PaginatedResultsMeta
}

// MovieLatest type is a struct for latest JSON response.
type MovieLatest struct {
	*MovieDetails
}

type Dates struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}

// MovieNowPlaying type is a struct for now playing JSON response.
type MovieNowPlaying struct {
	*MovieNowPlayingResults
	Dates Dates `json:"dates"`
	PaginatedResultsMeta
}

// MoviePopular type is a struct for popular JSON response.
type MoviePopular struct {
	*MoviePopularResults
	PaginatedResultsMeta
}

// MovieTopRated type is a struct for top rated JSON response.
type MovieTopRated struct {
	*MoviePopular
}

// MovieUpcoming type is a struct for upcoming JSON response.
type MovieUpcoming struct {
	*MovieNowPlaying
}

// MovieCredits type is a struct for credits JSON response.
type MovieCredits struct {
	ID   int64  `json:"id"`
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

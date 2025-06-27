package tmdb

// TVDetails type is a struct for details JSON response.
type TVDetails struct {
	VideoDetails
	TVShowMeta
	CreatedBy        []PersonCreditBase `json:"created_by"`
	EpisodeRunTime   []int              `json:"episode_run_time"`
	InProduction     bool               `json:"in_production"`
	Languages        []string           `json:"languages"`
	LastAirDate      string             `json:"last_air_date"`
	LastEpisodeToAir TVEpisode          `json:"last_episode_to_air"`
	NextEpisodeToAir *TVEpisode         `json:"next_episode_to_air"`
	Networks         []CompanyInfo      `json:"networks"`
	NumberOfEpisodes int                `json:"number_of_episodes"`
	NumberOfSeasons  int                `json:"number_of_seasons"`
	Seasons          []Season           `json:"seasons"`
	Type             string             `json:"type"`
	*TVAggregateCreditsAppend
	*TVAlternativeTitlesAppend
	*TVChangesAppend
	*TVContentRatingsAppend
	*TVCreditsAppend
	*TVEpisodeGroupsAppend
	*TVExternalIDsAppend
	*TVImagesAppend
	*TVKeywordsAppend
	*TVRecommendationsAppend
	*TVReviewsAppend
	*TVScreenedTheatricallyAppend
	*TVSimilarAppend
	*TVTranslationsAppend
	*TVVideosAppend
	*TVWatchProvidersAppend
}

// TVAggregateCreditsAppend type is a struct
// for aggregate credits in append to response.
type TVAggregateCreditsAppend struct {
	AggregateCredits TVAggregateCredits `json:"aggregate_credits"`
}

// TVAlternativeTitlesAppend type is a struct
// for alternative titles in append to response.
type TVAlternativeTitlesAppend struct {
	AlternativeTitles IDAlternativeTitleResults `json:"alternative_titles"`
}

// TVChangesAppend type is a struct for changes in append to response.
type TVChangesAppend struct {
	Changes Changes `json:"changes"`
}

// TVContentRatingsAppend type is a struct for
// content ratings in append to response.
type TVContentRatingsAppend struct {
	ContentRatings TVContentRatingsResults `json:"content_ratings"`
}

// TVCreditsAppend type is a struct for credits in append to response.
type TVCreditsAppend struct {
	Credits *TVCredits `json:"credits"`
}

// TVEpisodeGroupsAppend type is a struct for
// episode groups in append to response.
type TVEpisodeGroupsAppend struct {
	EpisodeGroups TVEpisodeGroupsResults `json:"episode_groups"`
}

// TVExternalIDsAppend type is a short for
// external ids in append to response.
type TVExternalIDsAppend struct {
	*TVExternalIDs `json:"external_ids"`
}

// TVImagesAppend type is a struct for images in append to response.
type TVImagesAppend struct {
	Images TVImages `json:"images"`
}

// TVKeywordsAppend type is a struct for keywords in append to response.
type TVKeywordsAppend struct {
	Keywords KeywordsResults `json:"keywords"`
}

// TVRecommendationsAppend type is a struct
// for recommendations in append to response.
type TVRecommendationsAppend struct {
	Recommendations PaginatedTVShowResults `json:"recommendations"`
}

// TVReviewsAppend type is a struct for reviews in append to response.
type TVReviewsAppend struct {
	Reviews PaginatedReviewsResults `json:"reviews"`
}

// TVScreenedTheatricallyAppend type is a struct
// for screened theatrically in append to response.
type TVScreenedTheatricallyAppend struct {
	ScreenedTheatrically *TVScreenedTheatrically `json:"screened_theatrically"`
}

// TVSimilarAppend type is a struct for
// similar tv shows in append to response.
type TVSimilarAppend struct {
	Similar *PaginatedTVShowResults `json:"similar"`
}

// TVTranslationsAppend type is a struct
// for translations in append to response.
type TVTranslationsAppend struct {
	Translations *TVTranslations `json:"translations"`
}

// TVVideosAppend type is a struct for videos in append to response.
type TVVideosAppend struct {
	Videos *VideoResults `json:"videos"`
}

// TVWatchProvidersAppend type is a struct for
// watch/providers in append to response.
type TVWatchProvidersAppend struct {
	WatchProviders IDWatchProviderResults `json:"watch/providers"`
}

// TVAccountStates type is a struct for account states JSON response.
type TVAccountStates struct {
	AccountStates
}

// TVAggregateCredits type is a struct for aggregate credits JSON response.
type TVAggregateCredits struct {
	Cast []CastCredit `json:"cast"`
	Crew []CrewCredit `json:"crew"`
}

type IDTVAggregateCredits struct {
	ID int64 `json:"id"`
	TVAggregateCredits
}

// TVExternalIDs type is a struct for external ids JSON response.
type TVExternalIDs struct {
	IMDbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
	WikidataID  string `json:"wikidata_id"`
	ID          int64  `json:"id,omitempty"`
}

// TVImages type is a struct for images JSON response.
type TVImages struct {
	Backdrops []ImageIso `json:"backdrops"`
	Logos     []ImageIso `json:"logos"`
	Posters   []ImageIso `json:"posters"`
}

type IDTVImages struct {
	ID int64 `json:"id"`
	TVImages
}

// TVScreenedTheatrically type is a struct for screened theatrically JSON response.
type TVScreenedTheatrically struct {
	ID int64 `json:"id"`
	*TVScreenedTheatricallyResults
}

type TVDataTranslation struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage"`
	Tagline  string `json:"tagline"`
}

type TVTranslation struct {
	TranslationBase
	Data TVDataTranslation `json:"data"`
}

// TVTranslations type is a struct for translations JSON response.
type TVTranslations struct {
	ID           int64           `json:"id"`
	Translations []TVTranslation `json:"translations"`
}

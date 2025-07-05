package tmdb

import json "github.com/goccy/go-json"

// TVDetails type is a struct for details JSON response.
type TVDetails struct {
	VideoDetails
	TVShowMeta
	CreatedBy        []CreatedBy `json:"created_by"`
	EpisodeRunTime   []int       `json:"episode_run_time"`
	InProduction     bool        `json:"in_production"`
	Languages        []string    `json:"languages"`
	LastAirDate      string      `json:"last_air_date"`
	LastEpisodeToAir TVEpisode   `json:"last_episode_to_air"`
	NextEpisodeToAir *TVEpisode  `json:"next_episode_to_air"`
	Networks         []Network   `json:"networks"`
	NumberOfEpisodes int         `json:"number_of_episodes"`
	NumberOfSeasons  int         `json:"number_of_seasons"`
	Seasons          []Season    `json:"seasons"`
	Type             string      `json:"type"`
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
	AggregateCredits *TVAggregateCredits `json:"aggregate_credits"`
}

// TVAlternativeTitlesAppend type is a struct
// for alternative titles in append to response.
type TVAlternativeTitlesAppend struct {
	AlternativeTitles *TVAlternativeTitles `json:"alternative_titles"`
}

// TVChangesAppend type is a struct for changes in append to response.
type TVChangesAppend struct {
	Changes *TVChanges `json:"changes"`
}

// TVContentRatingsAppend type is a struct for
// content ratings in append to response.
type TVContentRatingsAppend struct {
	ContentRatings *TVContentRatings `json:"content_ratings"`
}

// TVCreditsAppend type is a struct for credits in append to response.
type TVCreditsAppend struct {
	Credits *TVCredits `json:"credits"`
}

// TVEpisodeGroupsAppend type is a struct for
// episode groups in append to response.
type TVEpisodeGroupsAppend struct {
	EpisodeGroups *TVEpisodeGroups `json:"episode_groups"`
}

// TVExternalIDsAppend type is a short for
// external ids in append to response.
type TVExternalIDsAppend struct {
	*TVExternalIDs `json:"external_ids"`
}

// TVImagesAppend type is a struct for images in append to response.
type TVImagesAppend struct {
	Images *TVImages `json:"images"`
}

// TVKeywordsAppend type is a struct for keywords in append to response.
type TVKeywordsAppend struct {
	Keywords *TVKeywords `json:"keywords"`
}

// TVRecommendationsAppend type is a struct
// for recommendations in append to response.
type TVRecommendationsAppend struct {
	Recommendations *TVRecommendations `json:"recommendations"`
}

// TVReviewsAppend type is a struct for reviews in append to response.
type TVReviewsAppend struct {
	Reviews *TVReviews `json:"reviews"`
}

// TVScreenedTheatricallyAppend type is a struct
// for screened theatrically in append to response.
type TVScreenedTheatricallyAppend struct {
	ScreenedTheatrically *TVScreenedTheatrically `json:"screened_theatrically"`
}

// TVSimilarAppend type is a struct for
// similar tv shows in append to response.
type TVSimilarAppend struct {
	Similar *TVSimilar `json:"similar"`
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
	WatchProviders *WatchProviderResults `json:"watch/providers"`
}

// TVAccountStates type is a struct for account states JSON response.
type TVAccountStates struct {
	ID        int64           `json:"id"`
	Favorite  bool            `json:"favorite"`
	Rated     json.RawMessage `json:"rated"`
	Watchlist bool            `json:"watchlist"`
}

// TVAggregateCredits type is a struct for aggregate credits JSON response.
type TVAggregateCredits struct {
	ID   int64        `json:"id"`
	Cast []CastCredit `json:"cast"`
	Crew []CrewCredit `json:"crew"`
}

// TVAlternativeTitles type is a struct for alternative titles JSON response.
type TVAlternativeTitles struct {
	ID      int                `json:"id"`
	Results []AlternativeTitle `json:"results"`
}

// TVChanges type is a struct for changes JSON response.
type TVChanges struct {
	Changes []Change `json:"changes"`
}

// TVContentRatings type is a struct for content ratings JSON response.
type TVContentRatings struct {
	*TVContentRatingsResults
	ID int64 `json:"id"`
}

// TVCredits type is a struct for credits JSON response.
type TVCredits struct {
	ID   int64      `json:"id"`
	Cast []CastBase `json:"cast"`
	Crew []Crew     `json:"crew"`
}

// TVEpisodeGroups type is a struct for episode groups JSON response.
type TVEpisodeGroups struct {
	*TVEpisodeGroupsResults
	ID int64 `json:"id"`
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
	WikiDataID  string `json:"wikidata_id"`
	ID          int64  `json:"id,omitempty"`
}

// TVImage type is a struct for a single image.
type TVImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
}

// TVImages type is a struct for images JSON response.
type TVImages struct {
	ID        int64     `json:"id,omitempty"`
	Backdrops []TVImage `json:"backdrops"`
	Logos     []TVImage `json:"logos"`
	Posters   []TVImage `json:"posters"`
}

// TVKeywords type is a struct for keywords JSON response.
type TVKeywords struct {
	ID int64 `json:"id"`
	*TVKeywordsResults
}

// TVRecommendationsResults Result Types
type TVRecommendationsResults struct {
	Results []TVShow `json:"results"`
}

// TVRecommendations type is a struct for recommendations JSON response.
type TVRecommendations struct {
	*TVRecommendationsResults
	PaginatedResultsMeta
}

type TVRecommendationsMedia struct {
	*TVRecommendationsMediaResults
	PaginatedResultsMeta
}

// TVReviews type is a struct for reviews JSON response.
type TVReviews struct {
	ID int64 `json:"id"`
	*TVReviewsResults
	PaginatedResultsMeta
}

// TVScreenedTheatrically type is a struct for screened theatrically JSON response.
type TVScreenedTheatrically struct {
	ID int64 `json:"id"`
	*TVScreenedTheatricallyResults
}

// TVSimilar type is a struct for similar tv shows JSON response.
type TVSimilar struct {
	*TVRecommendations
}

type TVTranslationData struct {
	Name     string `json:"name,omitempty"`
	Overview string `json:"overview,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Tagline  string `json:"tagline,omitempty"`
}

type TVTranslation struct {
	Iso3166_1   string            `json:"iso_3166_1"`
	Iso639_1    string            `json:"iso_639_1"`
	Name        string            `json:"name"`
	EnglishName string            `json:"english_name"`
	Data        TVTranslationData `json:"data"`
}

// TVTranslations type is a struct for translations JSON response.
type TVTranslations struct {
	ID           int64           `json:"id"`
	Translations []TVTranslation `json:"translations"`
}

// TVLatest type is a struct for latest JSON response.
type TVLatest struct {
	*TVDetails
}

// TVAiringToday type is a struct for airing today JSON response.
type TVAiringToday struct {
	PaginatedResultsMeta
	*TVAiringTodayResults
}

// TVOnTheAir type is a struct for on the air JSON response.
type TVOnTheAir struct {
	*TVAiringToday
}

// TVPopular type is a struct for popular JSON response.
type TVPopular struct {
	*TVAiringToday
}

// TVTopRated type is a struct for top rated JSON response.
type TVTopRated struct {
	*TVAiringToday
}

package tmdb

import json "github.com/goccy/go-json"

// TVDetails type is a struct for details JSON response.
type TVDetails struct {
	BackdropPath        string              `json:"backdrop_path"`
	CreatedBy           []CreatedBy         `json:"created_by"`
	EpisodeRunTime      []int               `json:"episode_run_time"`
	FirstAirDate        string              `json:"first_air_date"`
	Genres              []Genre             `json:"genres"`
	Homepage            string              `json:"homepage"`
	ID                  int64               `json:"id"`
	InProduction        bool                `json:"in_production"`
	Languages           []string            `json:"languages"`
	LastAirDate         string              `json:"last_air_date"`
	Name                string              `json:"name"`
	LastEpisodeToAir    LastEpisodeToAir    `json:"last_episode_to_air"`
	NextEpisodeToAir    NextEpisodeToAir    `json:"next_episode_to_air"`
	Networks            []Network           `json:"networks"`
	NumberOfEpisodes    int                 `json:"number_of_episodes"`
	NumberOfSeasons     int                 `json:"number_of_seasons"`
	OriginCountry       []string            `json:"origin_country"`
	OriginalLanguage    string              `json:"original_language"`
	OriginalName        string              `json:"original_name"`
	Overview            string              `json:"overview"`
	Popularity          float32             `json:"popularity"`
	PosterPath          string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []ProductionCountry `json:"production_countries"`
	Seasons             []Season            `json:"seasons"`
	Status              string              `json:"status"`
	Tagline             string              `json:"tagline"`
	Type                string              `json:"type"`
	VoteMetrics
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
	AggregateCredits *TVAggregateCredits `json:"aggregate_credits,omitempty"`
}

// TVAlternativeTitlesAppend type is a struct
// for alternative titles in append to response.
type TVAlternativeTitlesAppend struct {
	AlternativeTitles *TVAlternativeTitles `json:"alternative_titles,omitempty"`
}

// TVChangesAppend type is a struct for changes in append to response.
type TVChangesAppend struct {
	Changes *TVChanges `json:"changes,omitempty"`
}

// TVContentRatingsAppend type is a struct for
// content ratings in append to response.
type TVContentRatingsAppend struct {
	ContentRatings *TVContentRatings `json:"content_ratings,omitempty"`
}

// TVCreditsAppend type is a struct for credits in append to response.
type TVCreditsAppend struct {
	Credits struct {
		*TVCredits
	} `json:"credits,omitempty"`
}

// TVEpisodeGroupsAppend type is a struct for
// episode groups in append to response.
type TVEpisodeGroupsAppend struct {
	EpisodeGroups *TVEpisodeGroups `json:"episode_groups,omitempty"`
}

// TVExternalIDsAppend type is a short for
// external ids in append to response.
type TVExternalIDsAppend struct {
	*TVExternalIDs `json:"external_ids,omitempty"`
}

// TVImagesAppend type is a struct for images in append to response.
type TVImagesAppend struct {
	Images *TVImages `json:"images,omitempty"`
}

// TVKeywordsAppend type is a struct for keywords in append to response.
type TVKeywordsAppend struct {
	Keywords struct {
		*TVKeywords
	} `json:"keywords,omitempty"`
}

// TVRecommendationsAppend type is a struct
// for recommendations in append to response.
type TVRecommendationsAppend struct {
	Recommendations *TVRecommendations `json:"recommendations,omitempty"`
}

// TVReviewsAppend type is a struct for reviews in append to response.
type TVReviewsAppend struct {
	Reviews struct {
		*TVReviews
	} `json:"reviews,omitempty"`
}

// TVScreenedTheatricallyAppend type is a struct
// for screened theatrically in append to response.
type TVScreenedTheatricallyAppend struct {
	ScreenedTheatrically *TVScreenedTheatrically `json:"screened_theatrically,omitempty"`
}

// TVSimilarAppend type is a struct for
// similar tv shows in append to response.
type TVSimilarAppend struct {
	Similar *TVSimilar `json:"similar,omitempty"`
}

// TVTranslationsAppend type is a struct
// for translations in append to response.
type TVTranslationsAppend struct {
	Translations *TVTranslations `json:"translations,omitempty"`
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
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		ID                 int64   `json:"id"`
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		Order              int     `json:"order"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		Roles              []struct {
			Character    string `json:"character"`
			CreditID     string `json:"credit_id"`
			EpisodeCount int    `json:"episode_count"`
		} `json:"roles"`
		TotalEpisodeCount int `json:"total_episode_count"`
	} `json:"cast"`
	Crew []struct {
		ID         int64  `json:"id"`
		Adult      bool   `json:"adult"`
		Department string `json:"department"`
		Gender     int    `json:"gender"`
		Jobs       []struct {
			CreditID     string `json:"credit_id"`
			EpisodeCount int    `json:"episode_count"`
			Job          string `json:"job"`
		} `json:"jobs"`
		TotalEpisodeCount  int     `json:"total_episode_count"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
}

// TVAlternativeTitles type is a struct for alternative titles JSON response.
type TVAlternativeTitles struct {
	ID      int                `json:"id"`
	Results []AlternativeTitle `json:"results"`
}

// TVChanges type is a struct for changes JSON response.
type TVChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string          `json:"id"`
			Action        string          `json:"action"`
			Time          string          `json:"time"`
			Iso639_1      string          `json:"iso_639_1"`
			Iso3166_1     string          `json:"iso_3166_1"`
			Value         json.RawMessage `json:"value"`
			OriginalValue json.RawMessage `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// TVContentRatings type is a struct for content ratings JSON response.
type TVContentRatings struct {
	*TVContentRatingsResults
	ID int64 `json:"id,omitempty"`
}

// TVCredits type is a struct for credits JSON response.
type TVCredits struct {
	ID   int64 `json:"id,omitempty"`
	Cast []struct {
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		Order              int     `json:"order"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		CreditID           string  `json:"credit_id"`
		Department         string  `json:"department"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		Job                string  `json:"job"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
}

// TVEpisodeGroups type is a struct for episode groups JSON response.
type TVEpisodeGroups struct {
	*TVEpisodeGroupsResults
	ID int64 `json:"id,omitempty"`
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
	WikiDataID  string `json:"wikidata_id,omitempty"`
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
	ID int64 `json:"id,omitempty"`
	*TVKeywordsResults
}

// TVRecommendations type is a struct for recommendations JSON response.
type TVRecommendations struct {
	*TVRecommendationsResults
	PaginatedResultsMeta
}

// TVReviews type is a struct for reviews JSON response.
type TVReviews struct {
	ID int64 `json:"id,omitempty"`
	*TVReviewsResults
	PaginatedResultsMeta
}

// TVScreenedTheatrically type is a struct for screened theatrically JSON response.
type TVScreenedTheatrically struct {
	ID int64 `json:"id,omitempty"`
	*TVScreenedTheatricallyResults
}

// TVSimilar type is a struct for similar tv shows JSON response.
type TVSimilar struct {
	*TVRecommendations
}

// TVTranslations type is a struct for translations JSON response.
type TVTranslations struct {
	ID           int64         `json:"id,omitempty"`
	Translations []Translation `json:"translations"`
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

package tmdb

// TVSeasonDetails is a struct for details JSON response.
type TVSeasonDetails struct {
	TVSeason
	IDString string                 `json:"_id"`
	Episodes []TVEpisodeDetailsBase `json:"episodes"`
	Networks []CompanyInfo          `json:"networks"`
	*TVAggregateCreditsAppend
	*TVSeasonCreditsAppend
	*TVSeasonExternalIDsAppend
	*TVSeasonImagesAppend
	*TVSeasonVideosAppend
	*TVSeasonTranslationsAppend
}

// TVSeasonCreditsAppend type is a struct
// for credits in append to response.
type TVSeasonCreditsAppend struct {
	Credits TVCredits `json:"credits"`
}

// TVSeasonExternalIDsAppend type is a struct
// for external ids in append to response.
type TVSeasonExternalIDsAppend struct {
	ExternalIDs TVSeasonExternalIDs `json:"external_ids"`
}

// TVSeasonImagesAppend type is a struct
// for images in append to response.
type TVSeasonImagesAppend struct {
	Images TVSeasonImages `json:"images"`
}

// TVSeasonTranslationsAppend type is a struct
// for translations in append to response.
type TVSeasonTranslationsAppend struct {
	Translations TVSeasonTranslations `json:"translations"`
}

// TVSeasonTranslations type is a struct
type TVSeasonTranslations struct {
	ID           int64                 `json:"id"`
	Translations []TVSeasonTranslation `json:"translations"`
}

type TVSeasonTranslation struct {
	TranslationBase
	Data TVSeasonDataTranslation `json:"data"`
}

type TVSeasonDataTranslation struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

// TVSeasonVideosAppend type is a struct
// for videos in append to response.
type TVSeasonVideosAppend struct {
	Videos *VideoResults `json:"videos"`
}

// TVSeasonExternalIDs type is a struct for external ids JSON response.
type TVSeasonExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
}

// TVSeasonImages type is a struct for images JSON response.
type TVSeasonImages struct {
	ID      int64      `json:"id"`
	Posters []ImageIso `json:"posters"`
}

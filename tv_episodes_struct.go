package tmdb

// TVEpisodeDetails type is a struct for details JSON response.
type TVEpisodeDetails struct {
	TVEpisodeDetailsBase
	*TVEpisodeCreditsAppend
	*TVEpisodeExternalIDsAppend
	*TVEpisodeImagesAppend
	*TVEpisodeTranslationsAppend
	*TVEpisodeVideosAppend
}

// TVEpisodeCreditsAppend type is a struct
// for credits in append to response.
type TVEpisodeCreditsAppend struct {
	Credits *TVEpisodeCredits `json:"credits"`
}

// TVEpisodeExternalIDsAppend type is a struct
// for external ids in append to response.
type TVEpisodeExternalIDsAppend struct {
	ExternalIDs *TVEpisodeExternalIDs `json:"external_ids"`
}

// TVEpisodeImagesAppend type is a struct
// for images in append to response.
type TVEpisodeImagesAppend struct {
	Images *TVEpisodeImages `json:"images"`
}

// TVEpisodeTranslationsAppend type is a struct
// for translations in append to response.
type TVEpisodeTranslationsAppend struct {
	Translations *TVEpisodeTranslations `json:"translations"`
}

// TVEpisodeVideosAppend type is a struct
// for videos in append to response.
type TVEpisodeVideosAppend struct {
	Videos *VideoResults `json:"videos"`
}

// TVEpisodeCredits type is a struct for credits JSON response.
type TVEpisodeCredits struct {
	Cast       []CastBase `json:"cast"`
	Crew       []Crew     `json:"crew"`
	GuestStars []CastBase `json:"guest_stars"`
	ID         int64      `json:"id"`
}

// TVEpisodeExternalIDs type is a struct for external ids JSON response.
type TVEpisodeExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	IMDbID      string `json:"imdb_id"`
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
}

// TVEpisodeImages type is a struct for images JSON response.
type TVEpisodeImages struct {
	ID     int64      `json:"id"`
	Stills []ImageIso `json:"stills"`
}

// TVEpisodeTranslations type is a struct for translations JSON response.
type TVEpisodeTranslations struct {
	ID           int64                  `json:"id"`
	Translations []TVEpisodeTranslation `json:"translations"`
}

// TVEpisodeRate type is a struct for rate JSON response.
type TVEpisodeRate struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

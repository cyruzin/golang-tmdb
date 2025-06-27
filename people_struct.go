package tmdb

type People struct {
	PersonBase
	PersonMeta
	Birthday     *string  `json:"birthday"`
	Deathday     *string  `json:"deathday"`
	AlsoKnownAs  []string `json:"also_known_as"`
	Biography    string   `json:"biography"`
	PlaceOfBirth *string  `json:"place_of_birth"`
	IMDbID       *string  `json:"imdb_id"`
	Homepage     *string  `json:"homepage"`
}

// PersonDetails type is a struct for details JSON response.
type PersonDetails struct {
	People
	*PersonChangesAppend
	*PersonMovieCreditsAppend
	*PersonTVCreditsAppend
	*PersonCombinedCreditsAppend
	*PersonExternalIDsAppend
	*PersonImagesAppend
	*PersonTranslationsAppend
}

// PersonChangesAppend type is a struct
// for changes JSON in append to response.
type PersonChangesAppend struct {
	Changes Changes `json:"changes"`
}

// PersonTVCreditsAppend type is a struct
// for tv credits in append to response.
type PersonTVCreditsAppend struct {
	TVCredits PersonTVCredits `json:"tv_credits"`
}

// PersonMovieCreditsAppend type is a struct
// for movie credits in append to response.
type PersonMovieCreditsAppend struct {
	MovieCredits PersonMovieCredits `json:"movie_credits"`
}

// PersonCombinedCreditsAppend type is a struct
// for combined credits in append to response.
type PersonCombinedCreditsAppend struct {
	CombinedCredits PersonMediaCredits `json:"combined_credits"`
}

// PersonExternalIDsAppend type is a struct
// for external ids in append to response.
type PersonExternalIDsAppend struct {
	ExternalIDs PersonExternalIDs `json:"external_ids"`
}

// PersonImagesAppend type is a struct
// for images in append to response.
type PersonImagesAppend struct {
	Images PersonImages `json:"images"`
}

// PersonTranslationsAppend type is a struct
// for translations in append to response.
type PersonTranslationsAppend struct {
	Translations PersonTranslations `json:"translations"`
}

// PersonMovieCredits type is a struct for movie credits JSON response.
type PersonMovieCredits struct {
	Cast []PeopleMovieCreditCast `json:"cast"`
	Crew []PeopleMovieCreditCrew `json:"crew"`
	ID   int64                   `json:"id"`
}

type PersonTVCredits struct {
	Cast []PeopleTVCreditCast `json:"cast"`
	Crew []PeopleTVCreditCrew `json:"crew"`
	ID   int64                `json:"id"`
}

// PersonExternalIDs type is a struct for external ids JSON response.
type PersonExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	FreebaseMid string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	IMDbID      string `json:"imdb_id"`
	TvrageID    int64  `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TiktokID    string `json:"tiktok_id"`
	TwitterID   string `json:"twitter_id"`
	YoutubeID   string `json:"youtube_id"`
}

// PersonImages type is a struct for images JSON response.
type PersonImages struct {
	Profiles []ImageIso `json:"profiles"`
	ID       int        `json:"id"`
}

// PersonTranslations type is a struct for translations JSON response.
type PersonTranslations struct {
	ID           int64               `json:"id"`
	Translations []PersonTranslation `json:"translations"`
}

type PersonTranslation struct {
	TranslationBase
	Data PersonDataTranslation `json:"data"`
}

type PersonDataTranslation struct {
	Biography string `json:"biography"`
	Name      string `json:"name"`
	Primary   bool   `json:"primary"`
}

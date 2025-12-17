package tmdb

// TVSeasonDetails is a struct for details JSON response.
type TVSeasonDetails struct {
	IDString string `json:"_id"`
	AirDate  string `json:"air_date"`
	Episodes []struct {
		AirDate        string `json:"air_date"`
		EpisodeNumber  int    `json:"episode_number"`
		ID             int64  `json:"id"`
		Name           string `json:"name"`
		Overview       string `json:"overview"`
		ProductionCode string `json:"production_code"`
		Runtime        int    `json:"runtime"`
		SeasonNumber   int    `json:"season_number"`
		ShowID         int64  `json:"show_id"`
		StillPath      string `json:"still_path"`
		VoteMetrics
		Crew []struct {
			ID          int64  `json:"id"`
			CreditID    string `json:"credit_id"`
			Name        string `json:"name"`
			Department  string `json:"department"`
			Job         string `json:"job"`
			Gender      int    `json:"gender"`
			ProfilePath string `json:"profile_path"`
		} `json:"crew"`
		GuestStars []struct {
			ID          int64  `json:"id"`
			Name        string `json:"name"`
			CreditID    string `json:"credit_id"`
			Character   string `json:"character"`
			Order       int    `json:"order"`
			Gender      int    `json:"gender"`
			ProfilePath string `json:"profile_path"`
		} `json:"guest_stars"`
	} `json:"episodes"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	ID           int64   `json:"id"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float32 `json:"vote_average"`
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
	Credits *TVSeasonCredits `json:"credits,omitempty"`
}

// TVSeasonExternalIDsAppend type is a struct
// for external ids in append to response.
type TVSeasonExternalIDsAppend struct {
	ExternalIDs *TVSeasonExternalIDs `json:"external_ids,omitempty"`
}

// TVSeasonImagesAppend type is a struct
// for images in append to response.
type TVSeasonImagesAppend struct {
	Images *TVSeasonImages `json:"images,omitempty"`
}

// TVSeasonTranslationsAppend type is a struct
// for translations in append to response.
type TVSeasonTranslationsAppend struct {
	Translations *TVSeasonTranslations `json:"translations,omitempty"`
}

// TVSeasonTranslations type is a struct
type TVSeasonTranslations struct {
	ID           int64         `json:"id,omitempty"`
	Translations []Translation `json:"translations"`
}

// TVSeasonVideosAppend type is a struct
// for videos in append to response.
type TVSeasonVideosAppend struct {
	Videos *VideoResults `json:"videos"`
}

// TVSeasonChanges is a struct for changes JSON response.
type TVSeasonChanges struct {
	Changes []struct {
		Items []struct {
			ID        string `json:"id"`
			Action    string `json:"action"`
			Time      string `json:"time"`
			Iso639_1  string `json:"iso_639_1"`
			Iso3166_1 string `json:"iso_3166_1"`
			Value     struct {
				EpisodeID     int64 `json:"episode_id"`
				EpisodeNumber int   `json:"episode_number"`
			} `json:"value"`
		} `json:"items"`
		Key string `json:"key"`
	} `json:"changes"`
}

// TVSeasonCredits type is a struct for credits JSON response.
type TVSeasonCredits struct {
	Cast []struct {
		Character   string `json:"character"`
		CreditID    string `json:"credit_id"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Order       int    `json:"order"`
		ProfilePath string `json:"profile_path"`
	} `json:"cast"`
	Crew []struct {
		CreditID    string `json:"credit_id"`
		Department  string `json:"department"`
		Gender      int    `json:"gender"`
		ID          int64  `json:"id"`
		Job         string `json:"job"`
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
	} `json:"crew"`
	ID int `json:"id"`
}

// TVSeasonExternalIDs type is a struct for external ids JSON response.
type TVSeasonExternalIDs struct {
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	ID          int64  `json:"id,omitempty"`
}

// TVSeasonImage type is a struct for a single image.
type TVSeasonImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
}

// TVSeasonImages type is a struct for images JSON response.
type TVSeasonImages struct {
	ID      int64           `json:"id,omitempty"`
	Posters []TVSeasonImage `json:"posters"`
}

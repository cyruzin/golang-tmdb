package tmdb

import json "github.com/goccy/go-json"

// TVEpisodeDetails type is a struct for details JSON response.
type TVEpisodeDetails struct {
	AirDate string `json:"air_date"`
	Crew    []struct {
		Department         string  `json:"department"`
		Job                string  `json:"job"`
		CreditID           string  `json:"credit_id"`
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
	EpisodeNumber int    `json:"episode_number"`
	EpisodeType   string `json:"episode_type"`
	GuestStars    []struct {
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Order              int     `json:"order"`
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"guest_stars"`
	Name           string `json:"name"`
	Overview       string `json:"overview"`
	ID             int64  `json:"id"`
	ProductionCode string `json:"production_code"`
	Runtime        int    `json:"runtime"`
	SeasonNumber   int    `json:"season_number"`
	StillPath      string `json:"still_path"`
	VoteMetrics
	*TVEpisodeCreditsAppend
	*TVEpisodeExternalIDsAppend
	*TVEpisodeImagesAppend
	*TVEpisodeTranslationsAppend
	*TVEpisodeVideosAppend
}

// TVEpisodeCreditsAppend type is a struct
// for credits in append to response.
type TVEpisodeCreditsAppend struct {
	Credits *TVEpisodeCredits `json:"credits,omitempty"`
}

// TVEpisodeExternalIDsAppend type is a struct
// for external ids in append to response.
type TVEpisodeExternalIDsAppend struct {
	ExternalIDs *TVEpisodeExternalIDs `json:"external_ids,omitempty"`
}

// TVEpisodeImagesAppend type is a struct
// for images in append to response.
type TVEpisodeImagesAppend struct {
	Images *TVEpisodeImages `json:"images,omitempty"`
}

// TVEpisodeTranslationsAppend type is a struct
// for translations in append to response.
type TVEpisodeTranslationsAppend struct {
	Translations *TVEpisodeTranslations `json:"translations,omitempty"`
}

// TVEpisodeVideosAppend type is a struct
// for videos in append to response.
type TVEpisodeVideosAppend struct {
	Videos *VideoResults `json:"videos"`
}

// TVEpisodeChanges type is a struct for changes JSON response.
type TVEpisodeChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string `json:"id"`
			Action        string `json:"action"`
			Time          string `json:"time"`
			Iso639_1      string `json:"iso_639_1"`
			Iso3166_1     string `json:"iso_3166_1"`
			OriginalValue struct {
				PersonID  int64  `json:"person_id"`
				Character string `json:"character"`
				Order     int64  `json:"order"`
				CreditID  string `json:"credit_id"`
			} `json:"original_values,omitempty"`
			Value json.RawMessage `json:"value,omitempty"`
		} `json:"items"`
	} `json:"changes"`
}

// TVEpisodeCredits type is a struct for credits JSON response.
type TVEpisodeCredits struct {
	Cast []struct {
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Order              int     `json:"order"`
	} `json:"cast"`
	Crew []struct {
		Department         string  `json:"department"`
		Job                string  `json:"job"`
		CreditID           string  `json:"credit_id"`
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"crew"`
	GuestStars []struct {
		Character          string  `json:"character"`
		CreditID           string  `json:"credit_id"`
		Order              int     `json:"order"`
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
	} `json:"guest_stars"`
	ID int64 `json:"id"`
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

// TVEpisodeImage type is a struct for a single image.
type TVEpisodeImage struct {
	ImageBase
	Iso3166_1 string `json:"iso_3166_1"`
	Iso639_1  string `json:"iso_639_1"`
}

// TVEpisodeImages type is a struct for images JSON response.
type TVEpisodeImages struct {
	ID     int64            `json:"id,omitempty"`
	Stills []TVEpisodeImage `json:"stills"`
}

// TVEpisodeTranslations type is a struct for translations JSON response.
type TVEpisodeTranslations struct {
	ID           int64         `json:"id,omitempty"`
	Translations []Translation `json:"translations"`
}

// TVEpisodeRate type is a struct for rate JSON response.
type TVEpisodeRate struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

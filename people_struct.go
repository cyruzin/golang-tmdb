package tmdb

// PersonDetails type is a struct for details JSON response.
type PersonDetails struct {
	Birthday           string   `json:"birthday"`
	KnownForDepartment string   `json:"known_for_department"`
	Deathday           string   `json:"deathday"`
	ID                 int64    `json:"id"`
	Name               string   `json:"name"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Gender             int      `json:"gender"`
	Biography          string   `json:"biography"`
	Popularity         float32  `json:"popularity"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	ProfilePath        string   `json:"profile_path"`
	Adult              bool     `json:"adult"`
	IMDbID             string   `json:"imdb_id"`
	Homepage           string   `json:"homepage"`
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
	Changes *PersonChanges `json:"changes,omitempty"`
}

// PersonTVCreditsAppend type is a struct
// for tv credits in append to response.
type PersonTVCreditsAppend struct {
	TVCredits *PersonTVCredits `json:"tv_credits,omitempty"`
}

// PersonMovieCreditsAppend type is a struct
// for movie credits in append to response.
type PersonMovieCreditsAppend struct {
	MovieCredits *PersonMovieCredits `json:"movie_credits,omitempty"`
}

// PersonCombinedCreditsAppend type is a struct
// for combined credits in append to response.
type PersonCombinedCreditsAppend struct {
	CombinedCredits *PersonCombinedCredits `json:"combined_credits,omitempty"`
}

// PersonExternalIDsAppend type is a struct
// for external ids in append to response.
type PersonExternalIDsAppend struct {
	ExternalIDs *PersonExternalIDs `json:"external_ids,omitempty"`
}

// PersonImagesAppend type is a struct
// for images in append to response.
type PersonImagesAppend struct {
	Images *PersonImages `json:"images,omitempty"`
}

// PersonTranslationsAppend type is a struct
// for translations in append to response.
type PersonTranslationsAppend struct {
	Translations *PersonTranslations `json:"translations,omitempty"`
}

// PersonChanges type is a struct for changes JSON response.
type PersonChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string `json:"id"`
			Action        string `json:"action"`
			Time          string `json:"time"`
			Iso639_1      string `json:"iso_639_1"`
			Iso3166_1     string `json:"iso_3166_1"`
			OriginalValue string `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// PersonMovieCredits type is a struct for movie credits JSON response.
type PersonMovieCredits struct {
	Cast []struct {
		Character        string  `json:"character"`
		CreditID         string  `json:"credit_id"`
		PosterPath       string  `json:"poster_path"`
		ID               int64   `json:"id"`
		Order            int     `json:"order"`
		Video            bool    `json:"video"`
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Popularity       float32 `json:"popularity"`
		Title            string  `json:"title"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
		VoteMetrics
	} `json:"cast"`
	Crew []struct {
		ID               int64   `json:"id"`
		Department       string  `json:"department"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Job              string  `json:"job"`
		Overview         string  `json:"overview"`
		Video            bool    `json:"video"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		GenreIDs         []int64 `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		PosterPath       string  `json:"poster_path"`
		CreditID         string  `json:"credit_id"`
		VoteMetrics
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonTVCredits type is a struct for tv credits JSON response.
type PersonTVCredits struct {
	Cast []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		Name             string   `json:"name"`
		VoteMetrics
		Character          string `json:"character"`
		CreditID           string `json:"credit_id"`
		EpisodeCount       int    `json:"episode_count"`
		FirstCreditAirDate string `json:"first_credit_air_date"`
	} `json:"cast"`
	Crew []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		Name             string   `json:"name"`
		VoteMetrics
		CreditID           string `json:"credit_id"`
		Department         string `json:"department"`
		EpisodeCount       int    `json:"episode_count"`
		FirstCreditAirDate string `json:"first_credit_air_date"`
		Job                string `json:"job"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonCombinedCredits type is a struct for combined credits JSON response.
type PersonCombinedCredits struct {
	Cast []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		OriginalTitle    string   `json:"original_title"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date"`
		ReleaseDate      string   `json:"release_date"`
		Title            string   `json:"title"`
		Video            bool     `json:"video"`
		Name             string   `json:"name"`
		VoteMetrics
		Character          string `json:"character"`
		CreditID           string `json:"credit_id"`
		EpisodeCount       int    `json:"episode_count"`
		FirstCreditAirDate string `json:"first_credit_air_date"`
		Order              int    `json:"order"`
		MediaType          string `json:"media_type"`
	} `json:"cast"`
	Crew []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIDs         []int64  `json:"genre_ids"`
		ID               int64    `json:"id"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalTitle    string   `json:"original_title"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		ReleaseDate      string   `json:"release_date"`
		Title            string   `json:"title"`
		Video            bool     `json:"video"`
		FirstAirDate     string   `json:"first_air_date"`
		Name             string   `json:"name"`
		VoteMetrics
		CreditID           string `json:"credit_id"`
		Department         string `json:"department"`
		EpisodeCount       int    `json:"episode_count"`
		FirstCreditAirDate string `json:"first_credit_air_date"`
		Job                string `json:"job"`
		MediaType          string `json:"media_type"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
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

// PersonImage type is a struct for a single image.
type PersonImage struct {
	ImageBase
	Iso3166_1 string `json:"iso_3166_1"`
	Iso639_1  string `json:"iso_639_1"`
}

// PersonImages type is a struct for images JSON response.
type PersonImages struct {
	Profiles []PersonImage `json:"profiles"`
	ID       int           `json:"id,omitempty"`
}

// PersonTranslations type is a struct for translations JSON response.
type PersonTranslations struct {
	Translations []Translation `json:"translations"`
	ID           int64         `json:"id,omitempty"`
}

// PersonLatest type is a struct for latest JSON response.
type PersonLatest struct {
	Adult              bool     `json:"adult"`
	AlsoKnownAs        []string `json:"also_known_as"`
	Biography          string   `json:"biography"`
	Birthday           string   `json:"birthday"`
	Deathday           string   `json:"deathday"`
	Gender             int      `json:"gender"`
	Homepage           string   `json:"homepage"`
	ID                 int64    `json:"id"`
	IMDbID             string   `json:"imdb_id"`
	KnownForDepartment string   `json:"known_for_department"`
	Name               string   `json:"name"`
	PlaceOfBirth       string   `json:"place_of_birth"`
	Popularity         float32  `json:"popularity"`
	ProfilePath        string   `json:"profile_path"`
}

// PersonPopular type is a struct for popular JSON response.
type PersonPopular struct {
	PaginatedResultsMeta
	Results []struct {
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		ID                 int64   `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float32 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		KnownFor           []struct {
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			ID               int64   `json:"id"`
			Title            string  `json:"title"`
			OriginalTitle    string  `json:"original_title"`
			Name             string  `json:"name"`
			OriginalName     string  `json:"original_name"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
			MediaType        string  `json:"media_type"`
			OriginalLanguage string  `json:"original_language"`
			GenreIDs         []int64 `json:"genre_ids"`
			Popularity       float32 `json:"popularity"`
			ReleaseDate      string  `json:"release_date"`
			FirstAirDate     string  `json:"first_air_date"`
			Video            bool    `json:"video"`
			VoteMetrics
			OriginCountry []string `json:"origin_country"`
		} `json:"known_for"`
	} `json:"results"`
}

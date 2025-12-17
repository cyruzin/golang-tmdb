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
		CreditID         string   `json:"credit_id"`
		OriginalName     string   `json:"original_name"`
		ID               int64    `json:"id"`
		GenreIDs         []int64  `json:"genre_ids"`
		Character        string   `json:"character"`
		Name             string   `json:"name"`
		PosterPath       string   `json:"poster_path"`
		Popularity       float32  `json:"popularity"`
		EpisodeCount     int      `json:"episode_count"`
		OriginalLanguage string   `json:"original_language"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		VoteMetrics
	} `json:"cast"`
	Crew []struct {
		ID               int64    `json:"id"`
		Department       string   `json:"department"`
		OriginalLanguage string   `json:"original_language"`
		EpisodeCount     int      `json:"episode_count"`
		Job              string   `json:"job"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
		OriginalName     string   `json:"original_name"`
		GenreIDs         []int64  `json:"genre_ids"`
		Name             string   `json:"name"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		Popularity       float32  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		CreditID         string   `json:"credit_id"`
		VoteMetrics
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonCombinedCredits type is a struct for combined credits JSON response.
type PersonCombinedCredits struct {
	Cast []struct {
		ID               int64    `json:"id"`
		Character        string   `json:"character"`
		OriginalTitle    string   `json:"original_title"`
		Overview         string   `json:"overview"`
		Video            bool     `json:"video"`
		MediaType        string   `json:"media_type"`
		ReleaseDate      string   `json:"release_date"`
		Title            string   `json:"title"`
		Popularity       float32  `json:"popularity"`
		OriginalLanguage string   `json:"original_language"`
		GenreIDs         []int64  `json:"genre_ids"`
		BackdropPath     string   `json:"backdrop_path"`
		Adult            bool     `json:"adult"`
		PosterPath       string   `json:"poster_path"`
		CreditID         string   `json:"credit_id"`
		EpisodeCount     int      `json:"episode_count"`
		OriginCountry    []string `json:"origin_country"`
		OriginalName     string   `json:"original_name"`
		Name             string   `json:"name"`
		FirstAirDate     string   `json:"first_air_date"`
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
		MediaType        string  `json:"media_type"`
		PosterPath       string  `json:"poster_path"`
		BackdropPath     string  `json:"backdrop_path"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		GenreIDs         []int64 `json:"genre_ids"`
		Adult            bool    `json:"adult"`
		ReleaseDate      string  `json:"release_date"`
		CreditID         string  `json:"credit_id"`
		VoteMetrics
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonExternalIDs type is a struct for external ids JSON response.
type PersonExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	TwitterID   string `json:"twitter_id"`
	FacebookID  string `json:"facebook_id"`
	TvrageID    int64  `json:"tvrage_id"`
	InstagramID string `json:"instagram_id"`
	FreebaseMid string `json:"freebase_mid"`
	IMDbID      string `json:"imdb_id"`
	FreebaseID  string `json:"freebase_id"`
}

// PersonImage type is a struct for a single image.
type PersonImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
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
	Birthday     string   `json:"birthday"`
	Deathday     string   `json:"deathday"`
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	AlsoKnownAs  []string `json:"also_known_as"`
	Gender       int      `json:"gender"`
	Biography    string   `json:"biography"`
	Popularity   float32  `json:"popularity"`
	PlaceOfBirth string   `json:"place_of_birth"`
	ProfilePath  string   `json:"profile_path"`
	Adult        bool     `json:"adult"`
	IMDbID       string   `json:"imdb_id"`
	Homepage     string   `json:"homepage"`
}

// PersonPopular type is a struct for popular JSON response.
type PersonPopular struct {
	PaginatedResultsMeta
	Results []struct {
		Popularity  float32 `json:"popularity"`
		ID          int64   `json:"id"`
		ProfilePath string  `json:"profile_path"`
		Name        string  `json:"name"`
		KnownFor    []struct {
			OriginalTitle    string   `json:"original_title,omitempty"`
			OriginalName     string   `json:"original_name,omitempty"`
			ID               int64    `json:"id"`
			MediaType        string   `json:"media_type"`
			Name             string   `json:"name,omitempty"`
			Title            string   `json:"title,omitempty"`
			PosterPath       string   `json:"poster_path"`
			FirstAirDate     string   `json:"first_air_date,omitempty"`
			ReleaseDate      string   `json:"release_date,omitempty"`
			Popularity       float32  `json:"popularity"`
			GenreIDs         []int64  `json:"genre_ids"`
			OriginalLanguage string   `json:"original_language"`
			BackdropPath     string   `json:"backdrop_path"`
			Overview         string   `json:"overview"`
			OriginCountry    []string `json:"origin_country,omitempty"`
			VoteMetrics
		} `json:"known_for"`
		Adult bool `json:"adult"`
	} `json:"results"`
}

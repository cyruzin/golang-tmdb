package tmdb

import "fmt"

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
	*PersonTaggedImagesAppend
	*PersonTranslationsAppend
}

// PersonChanges type is a struct for changes JSON response.
type PersonChanges struct {
	Changes []struct {
		Key   string `json:"key"`
		Items []struct {
			ID            string `json:"id"`
			Action        string `json:"action"`
			Time          string `json:"time"`
			OriginalValue struct {
				Profile struct {
					FilePath string `json:"file_path"`
				} `json:"profile"`
			} `json:"original_value"`
		} `json:"items"`
	} `json:"changes"`
}

// PersonChangesAppend type is a struct
// for changes JSON in append to response.
type PersonChangesAppend struct {
	Changes *PersonChanges `json:"changes,omitempty"`
}

// PersonMovieCredits type is a struct for movie credits JSON response.
type PersonMovieCredits struct {
	Cast []struct {
		Character        string  `json:"character"`
		CreditID         string  `json:"credit_id"`
		PosterPath       string  `json:"poster_path"`
		ID               int64   `json:"id"`
		Video            bool    `json:"video"`
		VoteCount        int64   `json:"vote_count"`
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Popularity       float32 `json:"popularity"`
		Title            string  `json:"title"`
		VoteAverage      float32 `json:"vote_average"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"cast"`
	Crew []struct {
		ID               int64   `json:"id"`
		Department       string  `json:"department"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Job              string  `json:"job"`
		Overview         string  `json:"overview"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		ReleaseDate      string  `json:"release_date"`
		VoteAverage      float32 `json:"vote_average"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		GenreIDs         []int64 `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		PosterPath       string  `json:"poster_path"`
		CreditID         string  `json:"credit_id"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonMovieCreditsAppend type is a struct
// for movie credits in append to response.
type PersonMovieCreditsAppend struct {
	MovieCredits *PersonMovieCredits `json:"movie_credits,omitempty"`
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
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		Popularity       float32  `json:"popularity"`
		EpisodeCount     int      `json:"episode_count"`
		OriginalLanguage string   `json:"original_language"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country"`
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
		VoteCount        int64    `json:"vote_count"`
		VoteAverage      float32  `json:"vote_average"`
		PosterPath       string   `json:"poster_path"`
		CreditID         string   `json:"credit_id"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonTVCreditsAppend type is a struct
// for tv credits in append to response.
type PersonTVCreditsAppend struct {
	TVCredits *PersonTVCredits `json:"tv_credits,omitempty"`
}

// PersonCombinedCredits type is a struct for combined credits JSON response.
type PersonCombinedCredits struct {
	Cast []struct {
		ID               int64    `json:"id"`
		Character        string   `json:"character"`
		OriginalTitle    string   `json:"original_title"`
		Overview         string   `json:"overview"`
		VoteCount        int64    `json:"vote_count"`
		Video            bool     `json:"video"`
		MediaType        string   `json:"media_type"`
		ReleaseDate      string   `json:"release_date"`
		VoteAverage      float32  `json:"vote_average"`
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
	} `json:"cast"`
	Crew []struct {
		ID               int64   `json:"id"`
		Department       string  `json:"department"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Job              string  `json:"job"`
		Overview         string  `json:"overview"`
		VoteCount        int64   `json:"vote_count"`
		Video            bool    `json:"video"`
		MediaType        string  `json:"media_type"`
		PosterPath       string  `json:"poster_path"`
		BackdropPath     string  `json:"backdrop_path"`
		Title            string  `json:"title"`
		Popularity       float32 `json:"popularity"`
		GenreIDs         []int64 `json:"genre_ids"`
		VoteAverage      float32 `json:"vote_average"`
		Adult            bool    `json:"adult"`
		ReleaseDate      string  `json:"release_date"`
		CreditID         string  `json:"credit_id"`
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// PersonCombinedCreditsAppend type is a struct
// for combined credits in append to response.
type PersonCombinedCreditsAppend struct {
	CombinedCredits *PersonCombinedCredits `json:"combined_credits,omitempty"`
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

// PersonExternalIDsAppend type is a struct
// for external ids in append to response.
type PersonExternalIDsAppend struct {
	ExternalIDs *PersonExternalIDs `json:"external_ids,omitempty"`
}

// PersonImages type is a struct for images JSON response.
type PersonImages struct {
	Profiles []struct {
		Iso639_1    string  `json:"iso_639_1"`
		Width       int     `json:"width"`
		Height      int     `json:"height"`
		VoteCount   int64   `json:"vote_count"`
		VoteAverage float32 `json:"vote_average"`
		FilePath    string  `json:"file_path"`
		AspectRatio float32 `json:"aspect_ratio"`
	} `json:"profiles"`
	ID int `json:"id,omitempty"`
}

// PersonImagesAppend type is a struct
// for images in append to response.
type PersonImagesAppend struct {
	Images *PersonImages `json:"images,omitempty"`
}

// PersonTaggedImages type is a struct for tagged images JSON response.
type PersonTaggedImages struct {
	ID           int64 `json:"id"`
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
		Iso639_1    string  `json:"iso_639_1"`
		VoteCount   int64   `json:"vote_count"`
		MediaType   string  `json:"media_type"`
		FilePath    string  `json:"file_path"`
		AspectRatio float32 `json:"aspect_ratio"`
		Media       struct {
			Popularity       float32 `json:"popularity"`
			VoteCount        int64   `json:"vote_count"`
			Video            bool    `json:"video"`
			PosterPath       string  `json:"poster_path"`
			ID               int64   `json:"id"`
			Adult            bool    `json:"adult"`
			BackdropPath     string  `json:"backdrop_path"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			GenreIDs         []int64 `json:"genre_ids"`
			Title            string  `json:"title"`
			VoteAverage      float32 `json:"vote_average"`
			Overview         string  `json:"overview"`
			ReleaseDate      string  `json:"release_date"`
		} `json:"media"`
		Height      int     `json:"height"`
		VoteAverage float32 `json:"vote_average"`
		Width       int     `json:"width"`
	} `json:"results"`
}

// PersonTaggedImagesAppend type is a struct
// for tagged images in append to response.
type PersonTaggedImagesAppend struct {
	TaggedImages *PersonTaggedImages `json:"tagged_images,omitempty"`
}

// PersonTranslations type is a struct for translations JSON response.
type PersonTranslations struct {
	Translations []struct {
		Iso639_1  string `json:"iso_639_1"`
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
		Data      struct {
			Biography string `json:"biography"`
		} `json:"data"`
		EnglishName string `json:"english_name"`
	} `json:"translations"`
	ID int64 `json:"id,omitempty"`
}

// PersonTranslationsAppend type is a struct
// for translations in append to response.
type PersonTranslationsAppend struct {
	Translations *PersonTranslations `json:"translations,omitempty"`
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
	Page         int64 `json:"page"`
	TotalResults int64 `json:"total_results"`
	TotalPages   int64 `json:"total_pages"`
	Results      []struct {
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
			VoteCount        int64    `json:"vote_count"`
			VoteAverage      float32  `json:"vote_average"`
			PosterPath       string   `json:"poster_path"`
			FirstAirDate     string   `json:"first_air_date,omitempty"`
			ReleaseDate      string   `json:"release_date,omitempty"`
			Popularity       float32  `json:"popularity"`
			GenreIDs         []int64  `json:"genre_ids"`
			OriginalLanguage string   `json:"original_language"`
			BackdropPath     string   `json:"backdrop_path"`
			Overview         string   `json:"overview"`
			OriginCountry    []string `json:"origin_country,omitempty"`
		} `json:"known_for"`
		Adult bool `json:"adult"`
	} `json:"results"`
}

// GetPersonDetails get the primary person details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/people/get-person-details
func (c *Client) GetPersonDetails(
	id int, o map[string]string,
) (*PersonDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonDetails{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonChanges get the changes for a person.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by
// using the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/people/get-person-changes
func (c *Client) GetPersonChanges(
	id int, o map[string]string,
) (*PersonChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonChanges{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonMovieCredits get the movie credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-movie-credits
func (c *Client) GetPersonMovieCredits(
	id int, o map[string]string,
) (*PersonMovieCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/movie_credits?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonMovieCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonTVCredits get the TV show credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-tv-credits
func (c *Client) GetPersonTVCredits(
	id int, o map[string]string,
) (*PersonTVCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tv_credits?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonTVCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonCombinedCredits get the movie and TV credits together in a single response.
//
// https://developers.themoviedb.org/3/people/get-person-combined-credits
func (c *Client) GetPersonCombinedCredits(
	id int, o map[string]string,
) (*PersonCombinedCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/combined_credits?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonCombinedCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonExternalIDs get the external ids for a person.
// We currently support the following external sources.
//
// External Sources: IMDb ID, Facebook, Freebase MID, Freebase ID,
// Instagram, TVRage ID, Twitter.
//
// https://developers.themoviedb.org/3/people/get-person-external-ids
func (c *Client) GetPersonExternalIDs(
	id int, o map[string]string,
) (*PersonExternalIDs, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonExternalIDs{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonImages get the images for a person.
//
// https://developers.themoviedb.org/3/people/get-person-images
func (c *Client) GetPersonImages(id int) (*PersonImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s",
		baseURL, personURL, id, c.APIKey,
	)
	p := PersonImages{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonTaggedImages get the images that this person has been tagged in.
//
// https://developers.themoviedb.org/3/people/get-tagged-images
func (c *Client) GetPersonTaggedImages(
	id int, o map[string]string,
) (*PersonTaggedImages, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tagged_images?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonTaggedImages{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonTranslations get the images that this person has been tagged in.
//
// https://developers.themoviedb.org/3/people/get-tagged-images
func (c *Client) GetPersonTranslations(
	id int, o map[string]string,
) (*PersonTranslations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL, personURL, id, c.APIKey, options,
	)
	p := PersonTranslations{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonLatest get the most newly created person.
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/people/get-latest-person
func (c *Client) GetPersonLatest(
	o map[string]string,
) (*PersonLatest, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s",
		baseURL, personURL, c.APIKey, options,
	)
	p := PersonLatest{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPersonPopular get the list of popular people on TMDb.
// This list updates daily.
//
// https://developers.themoviedb.org/3/people/get-popular-people
func (c *Client) GetPersonPopular(
	o map[string]string,
) (*PersonPopular, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL, personURL, c.APIKey, options,
	)
	p := PersonPopular{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// TODO: Check Popular func.

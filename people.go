package tmdb

import "fmt"

// PeopleDetails type is a struct for details JSON response.
type PeopleDetails struct {
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
	*PeopleChangesShort
	*PeopleMovieCreditsShort
	*PeopleTVCreditsShort
	*PeopleCombinedCreditsShort
	*PeopleExternalIDsShort
	*PeopleImagesShort
	*PeopleTaggedImagesShort
	*PeopleTranslationsShort
}

// PeopleChanges type is a struct for changes JSON response.
type PeopleChanges struct {
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

// PeopleChangesShort type is a short struct for changes JSON response.
type PeopleChangesShort struct {
	Changes *PeopleChanges `json:"changes,omitempty"`
}

// PeopleMovieCredits type is a struct for movie credits JSON response.
type PeopleMovieCredits struct {
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

// PeopleMovieCreditsShort type is a short struct for movie credits JSON response.
type PeopleMovieCreditsShort struct {
	MovieCredits *PeopleMovieCredits `json:"movie_credits,omitempty"`
}

// PeopleTVCredits type is a struct for tv credits JSON response.
type PeopleTVCredits struct {
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

// PeopleTVCreditsShort type is a short struct for tv credits JSON response.
type PeopleTVCreditsShort struct {
	TVCredits *PeopleTVCredits `json:"tv_credits,omitempty"`
}

// PeopleCombinedCredits type is a struct for combined credits JSON response.
type PeopleCombinedCredits struct {
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

// PeopleCombinedCreditsShort type is a short struct for combined credits JSON response.
type PeopleCombinedCreditsShort struct {
	CombinedCredits *PeopleCombinedCredits `json:"combined_credits,omitempty"`
}

// PeopleExternalIDs type is a struct for external ids JSON response.
type PeopleExternalIDs struct {
	ID          int64  `json:"id,omitempty"`
	TwitterID   string `json:"twitter_id"`
	FacebookID  string `json:"facebook_id"`
	TvrageID    int64  `json:"tvrage_id"`
	InstagramID string `json:"instagram_id"`
	FreebaseMid string `json:"freebase_mid"`
	IMDbID      string `json:"imdb_id"`
	FreebaseID  string `json:"freebase_id"`
}

// PeopleExternalIDsShort type is a short struct for external ids JSON response.
type PeopleExternalIDsShort struct {
	ExternalIDs *PeopleExternalIDs `json:"external_ids,omitempty"`
}

// PeopleImages type is a struct for images JSON response.
type PeopleImages struct {
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

// PeopleImagesShort type is a short struct for images JSON response.
type PeopleImagesShort struct {
	Images *PeopleImages `json:"images,omitempty"`
}

// PeopleTaggedImages type is a struct for tagged images JSON response.
type PeopleTaggedImages struct {
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

// PeopleTaggedImagesShort type is a short struct for tagged images JSON response.
type PeopleTaggedImagesShort struct {
	TaggedImages *PeopleTaggedImages `json:"tagged_images,omitempty"`
}

// PeopleTranslations type is a struct for translations JSON response.
type PeopleTranslations struct {
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

// PeopleTranslationsShort type is a short struct for translations JSON response.
type PeopleTranslationsShort struct {
	Translations *PeopleTranslations `json:"translations,omitempty"`
}

// PeopleLatest type is a struct for latest JSON response.
type PeopleLatest struct {
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

// PeoplePopular type is a struct for popular JSON response.
type PeoplePopular struct {
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

// GetPeopleDetails get the primary person details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/people/get-person-details
func (c *Client) GetPeopleDetails(id int, o map[string]string) (*PeopleDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleDetails{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleChanges get the changes for a person.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by
// using the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/people/get-person-changes
func (c *Client) GetPeopleChanges(id int, o map[string]string) (*PeopleChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleChanges{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleMovieCredits get the movie credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-movie-credits
func (c *Client) GetPeopleMovieCredits(id int, o map[string]string) (*PeopleMovieCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/movie_credits?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleMovieCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleTVCredits get the TV show credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-tv-credits
func (c *Client) GetPeopleTVCredits(id int, o map[string]string) (*PeopleTVCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tv_credits?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleTVCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleCombinedCredits get the movie and TV credits together in a single response.
//
// https://developers.themoviedb.org/3/people/get-person-combined-credits
func (c *Client) GetPeopleCombinedCredits(id int, o map[string]string) (*PeopleCombinedCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/combined_credits?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleCombinedCredits{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleExternalIDs get the external ids for a person.
// We currently support the following external sources.
//
// External Sources: IMDb ID, Facebook, Freebase MID, Freebase ID,
// Instagram, TVRage ID, Twitter.
//
// https://developers.themoviedb.org/3/people/get-person-external-ids
func (c *Client) GetPeopleExternalIDs(id int, o map[string]string) (*PeopleExternalIDs, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleExternalIDs{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleImages get the images for a person.
//
// https://developers.themoviedb.org/3/people/get-person-images
func (c *Client) GetPeopleImages(id int) (*PeopleImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s", baseURL, personURL, id, c.APIKey,
	)
	p := PeopleImages{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleTaggedImages get the images that this person has been tagged in.
//
// https://developers.themoviedb.org/3/people/get-tagged-images
func (c *Client) GetPeopleTaggedImages(id int, o map[string]string) (*PeopleTaggedImages, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tagged_images?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleTaggedImages{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleTranslations get the images that this person has been tagged in.
//
// https://developers.themoviedb.org/3/people/get-tagged-images
func (c *Client) GetPeopleTranslations(id int, o map[string]string) (*PeopleTranslations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s", baseURL, personURL, id, c.APIKey, options,
	)
	p := PeopleTranslations{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeopleLatest get the most newly created person.
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/people/get-latest-person
func (c *Client) GetPeopleLatest(o map[string]string) (*PeopleLatest, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s", baseURL, personURL, c.APIKey, options,
	)
	p := PeopleLatest{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPeoplePopular get the list of popular people on TMDb.
// This list updates daily.
//
// https://developers.themoviedb.org/3/people/get-popular-people
func (c *Client) GetPeoplePopular(o map[string]string) (*PeoplePopular, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s", baseURL, personURL, c.APIKey, options,
	)
	p := PeoplePopular{}
	err := c.get(tmdbURL, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// TODO: Check Popular func.

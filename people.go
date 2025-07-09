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

// PersonTaggedImagesAppend type is a struct
// for tagged images in append to response.
type PersonTaggedImagesAppend struct {
	TaggedImages *PaginatedPersonTaggedImagesResults `json:"tagged_images"`
}

// PersonTranslationsAppend type is a struct
// for translations in append to response.
type PersonTranslationsAppend struct {
	Translations *PersonTranslations `json:"translations,omitempty"`
}

type IDPaginatedPersonTaggedImagesResults struct {
	ID int64 `json:"id"`
	PaginatedPersonTaggedImagesResults
}

// GetPersonDetails get the primary person details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/people/get-person-details
func (c *Client) GetPersonDetails(
	id int,
	urlOptions map[string]string,
) (*PersonDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personDetails := PersonDetails{}
	if err := c.get(tmdbURL, &personDetails); err != nil {
		return nil, err
	}
	return &personDetails, nil
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

// GetPersonChanges get the changes for a person.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by
// using the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/people/get-person-changes
func (c *Client) GetPersonChanges(
	id int,
	urlOptions map[string]string,
) (*PersonChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personChanges := PersonChanges{}
	err := c.get(tmdbURL, &personChanges)
	if err != nil {
		return nil, err
	}
	return &personChanges, nil
}

// PersonMovieCredits type is a struct for movie credits JSON response.
type PersonMovieCredits struct {
	Cast []PeopleMovieCreditCast `json:"cast"`
	Crew []PeopleMovieCreditCrew `json:"crew"`
	ID   int64                   `json:"id"`
}

// GetPersonMovieCredits get the movie credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-movie-credits
func (c *Client) GetPersonMovieCredits(
	id int,
	urlOptions map[string]string,
) (*PersonMovieCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/movie_credits?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personMovieCredits := PersonMovieCredits{}
	if err := c.get(tmdbURL, &personMovieCredits); err != nil {
		return nil, err
	}
	return &personMovieCredits, nil
}

// PersonTVCredits type is a struct for tv credits JSON response.
type PersonTVCredits struct {
	Cast []PeopleTVCreditCast `json:"cast"`
	Crew []PeopleTVCreditCrew `json:"crew"`
	ID   int64                `json:"id,omitempty"`
}

// GetPersonTVCredits get the TV show credits for a person.
//
// https://developers.themoviedb.org/3/people/get-person-tv-credits
func (c *Client) GetPersonTVCredits(
	id int,
	urlOptions map[string]string,
) (*PersonTVCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tv_credits?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personTVCredits := PersonTVCredits{}
	if err := c.get(tmdbURL, &personTVCredits); err != nil {
		return nil, err
	}
	return &personTVCredits, nil
}

// PersonCombinedCredits type is a struct for combined credits JSON response.
type PersonCombinedCredits struct {
	Cast []struct {
		ID               int64     `json:"id"`
		Character        string    `json:"character"`
		OriginalTitle    string    `json:"original_title"`
		Overview         string    `json:"overview"`
		Video            bool      `json:"video"`
		MediaType        MediaType `json:"media_type"`
		ReleaseDate      string    `json:"release_date"`
		Title            string    `json:"title"`
		Popularity       float32   `json:"popularity"`
		OriginalLanguage string    `json:"original_language"`
		GenreIDs         []int64   `json:"genre_ids"`
		BackdropPath     string    `json:"backdrop_path"`
		Adult            bool      `json:"adult"`
		PosterPath       string    `json:"poster_path"`
		CreditID         string    `json:"credit_id"`
		EpisodeCount     int       `json:"episode_count"`
		OriginCountry    []string  `json:"origin_country"`
		OriginalName     string    `json:"original_name"`
		Name             string    `json:"name"`
		FirstAirDate     string    `json:"first_air_date"`
		VoteMetrics
	} `json:"cast"`
	Crew []struct {
		ID               int64     `json:"id"`
		Department       string    `json:"department"`
		OriginalLanguage string    `json:"original_language"`
		OriginalTitle    string    `json:"original_title"`
		Job              string    `json:"job"`
		Overview         string    `json:"overview"`
		Video            bool      `json:"video"`
		MediaType        MediaType `json:"media_type"`
		PosterPath       string    `json:"poster_path"`
		BackdropPath     string    `json:"backdrop_path"`
		Title            string    `json:"title"`
		Popularity       float32   `json:"popularity"`
		GenreIDs         []int64   `json:"genre_ids"`
		Adult            bool      `json:"adult"`
		ReleaseDate      string    `json:"release_date"`
		CreditID         string    `json:"credit_id"`
		VoteMetrics
	} `json:"crew"`
	ID int64 `json:"id,omitempty"`
}

// GetPersonCombinedCredits get the movie and TV credits together in a single response.
//
// https://developers.themoviedb.org/3/people/get-person-combined-credits
func (c *Client) GetPersonCombinedCredits(
	id int,
	urlOptions map[string]string,
) (*PersonCombinedCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/combined_credits?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personCombinedCredits := PersonCombinedCredits{}
	if err := c.get(tmdbURL, &personCombinedCredits); err != nil {
		return nil, err
	}
	return &personCombinedCredits, nil
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

// GetPersonExternalIDs get the external ids for a person.
// We currently support the following external sources.
//
// External Sources: IMDb ID, Facebook, Freebase MID, Freebase ID,
// Instagram, TVRage ID, Twitter.
//
// https://developers.themoviedb.org/3/people/get-person-external-ids
func (c *Client) GetPersonExternalIDs(
	id int,
	urlOptions map[string]string,
) (*PersonExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personExternalIDS := PersonExternalIDs{}
	if err := c.get(tmdbURL, &personExternalIDS); err != nil {
		return nil, err
	}
	return &personExternalIDS, nil
}

// PersonImages type is a struct for images JSON response.
type PersonImages struct {
	Profiles []ImageIso `json:"profiles"`
	ID       int        `json:"id"`
}

// GetPersonImages get the images for a person.
//
// https://developer.themoviedb.org/reference/person-images
func (c *Client) GetPersonImages(
	id int,
) (*PersonImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
	)
	personImages := PersonImages{}
	if err := c.get(tmdbURL, &personImages); err != nil {
		return nil, err
	}
	return &personImages, nil
}

// Deprecated: Use GetPersonImages instead.
//
// GetPersonTaggedImages get the images that this person has been tagged in.
//
// This endpoint is deprecated in the TMDB API.
//
// https://developer.themoviedb.org/reference/person-tagged-images
func (c *Client) GetPersonTaggedImages(
	id int,
	urlOptions map[string]string,
) (*IDPaginatedPersonTaggedImagesResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/tagged_images?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personTaggedImages := IDPaginatedPersonTaggedImagesResults{}
	if err := c.get(tmdbURL, &personTaggedImages); err != nil {
		return nil, err
	}
	return &personTaggedImages, nil
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

// GetPersonTranslations get a list of translations that have been created for a person.
//
// https://developer.themoviedb.org/reference/translations
func (c *Client) GetPersonTranslations(
	id int,
	urlOptions map[string]string,
) (*PersonTranslations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL,
		personURL,
		id,
		c.apiKey,
		options,
	)
	personTranslations := PersonTranslations{}
	if err := c.get(tmdbURL, &personTranslations); err != nil {
		return nil, err
	}
	return &personTranslations, nil
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

// GetPersonLatest get the most newly created person.
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/people/get-latest-person
func (c *Client) GetPersonLatest(
	urlOptions map[string]string,
) (*PersonLatest, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s",
		baseURL,
		personURL,
		c.apiKey,
		options,
	)
	personLatest := PersonLatest{}
	if err := c.get(tmdbURL, &personLatest); err != nil {
		return nil, err
	}
	return &personLatest, nil
}

// GetPersonPopular get the list of popular people on TMDb.
// This list updates daily.
//
// https://developers.themoviedb.org/3/people/get-popular-people
func (c *Client) GetPersonPopular(
	urlOptions map[string]string,
) (*PaginatedPersonMediaResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL,
		personURL,
		c.apiKey,
		options,
	)
	personPopular := PaginatedPersonMediaResults{}
	if err := c.get(tmdbURL, &personPopular); err != nil {
		return nil, err
	}
	return &personPopular, nil
}

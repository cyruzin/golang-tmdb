package tmdb

import "fmt"

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

// GetTVSeasonDetails get the TV season details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
func (c *Client) GetTVSeasonDetails(
	id int,
	seasonNumber int,
	urlOptions map[string]string,
) (*TVSeasonDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonDetails := TVSeasonDetails{}
	if err := c.get(tmdbURL, &tvSeasonDetails); err != nil {
		return nil, err
	}
	return &tvSeasonDetails, nil
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

// GetTVSeasonChanges get the changes for a TV season.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-changes
func (c *Client) GetTVSeasonChanges(
	id int,
	urlOptions map[string]string,
) (*TVSeasonChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sseason/%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvSeasonChanges := TVSeasonChanges{}
	if err := c.get(tmdbURL, &tvSeasonChanges); err != nil {
		return nil, err
	}
	return &tvSeasonChanges, nil
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

// GetTVSeasonCredits get the credits for TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-credits
func (c *Client) GetTVSeasonCredits(
	id int,
	seasonNumber int,
	urlOptions map[string]string,
) (*TVSeasonCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/credits?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tVSeasonCredits := TVSeasonCredits{}
	if err := c.get(tmdbURL, &tVSeasonCredits); err != nil {
		return nil, err
	}
	return &tVSeasonCredits, nil
}

// TVSeasonExternalIDs type is a struct for external ids JSON response.
type TVSeasonExternalIDs struct {
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	ID          int64  `json:"id,omitempty"`
}

// GetTVSeasonExternalIDs get the external ids for a TV season.
// We currently support the following external sources.
//
// Media Databases: TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-external-ids
func (c *Client) GetTVSeasonExternalIDs(
	id int,
	seasonNumber int,
	urlOptions map[string]string,
) (*TVSeasonExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/external_ids?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonExternalIDs := TVSeasonExternalIDs{}
	if err := c.get(tmdbURL, &tvSeasonExternalIDs); err != nil {
		return nil, err
	}
	return &tvSeasonExternalIDs, nil
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

// GetTVSeasonImages get the images that belong to a TV season.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-images
func (c *Client) GetTVSeasonImages(
	id int,
	seasonNumber int,
	urlOptions map[string]string,
) (*TVSeasonImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/images?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonImages := TVSeasonImages{}
	if err := c.get(tmdbURL, &tvSeasonImages); err != nil {
		return nil, err
	}
	return &tvSeasonImages, nil
}

// GetTVSeasonVideos get the videos that have been added to a TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-videos
func (c *Client) GetTVSeasonVideos(
	id int,
	seasonNumber int,
	urlOptions map[string]string,
) (*VideoResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/videos?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
		options,
	)
	tvSeasonVideos := VideoResults{}
	if err := c.get(tmdbURL, &tvSeasonVideos); err != nil {
		return nil, err
	}
	return &tvSeasonVideos, nil
}

// GetTVSeasonTranslations get the translation data for an season.
//
// https://developer.themoviedb.org/reference/tv-season-translations
func (c *Client) GetTVSeasonTranslations(
	id int,
	seasonNumber int,
) (*TVSeasonTranslations, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/translations?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		c.apiKey,
	)
	tvSeasonTranslations := TVSeasonTranslations{}
	if err := c.get(tmdbURL, &tvSeasonTranslations); err != nil {
		return nil, err
	}
	return &tvSeasonTranslations, nil
}

package tmdb

import "fmt"

// TVSeasonsDetails is a struct for details JSON response.
type TVSeasonsDetails struct {
	IDString string `json:"_id"`
	AirDate  string `json:"air_date"`
	Episodes []struct {
		AirDate        string  `json:"air_date"`
		EpisodeNumber  int     `json:"episode_number"`
		ID             int64   `json:"id"`
		Name           string  `json:"name"`
		Overview       string  `json:"overview"`
		ProductionCode string  `json:"production_code"`
		SeasonNumber   int     `json:"season_number"`
		ShowID         int64   `json:"show_id"`
		StillPath      string  `json:"still_path"`
		VoteAverage    float32 `json:"vote_average"`
		VoteCount      int64   `json:"vote_count"`
		Crew           []struct {
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
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	ID           int64  `json:"id"`
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
	*TVSeasonsCreditsAppend
	*TVSeasonsExternalIDsAppend
	*TVSeasonsImagesAppend
	*TVSeasonsVideosAppend
}

// TVSeasonsChanges is a struct for changes JSON response.
type TVSeasonsChanges struct {
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

// TVSeasonsCredits type is a struct for credits JSON response.
type TVSeasonsCredits struct {
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

// TVSeasonsCreditsAppend type is a struct
// for credits in append to response.
type TVSeasonsCreditsAppend struct {
	Credits *TVSeasonsCredits `json:"credits,omitempty"`
}

// TVSeasonsExternalIDs type is a struct for external ids JSON response.
type TVSeasonsExternalIDs struct {
	FreebaseMID string `json:"freebase_mid"`
	FreebaseID  string `json:"freebase_id"`
	TVDBID      int64  `json:"tvdb_id"`
	TVRageID    int64  `json:"tvrage_id"`
	ID          int64  `json:"id,omitempty"`
}

// TVSeasonsExternalIDsAppend type is a struct
// for external ids in append to response.
type TVSeasonsExternalIDsAppend struct {
	ExternalIDs *TVSeasonsExternalIDs `json:"external_ids,omitempty"`
}

// TVSeasonsImages type is a struct for images JSON response.
type TVSeasonsImages struct {
	ID      int64 `json:"id,omitempty"`
	Posters []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"posters"`
}

// TVSeasonsImagesAppend type is a struct
// for images in append to response.
type TVSeasonsImagesAppend struct {
	Images *TVSeasonsImages `json:"images,omitempty"`
}

// TVSeasonsVideos type is a struct for videos JSON response.
type TVSeasonsVideos struct {
	ID      int64 `json:"id,omitempty"`
	Results []struct {
		ID        string `json:"id"`
		Iso639_1  string `json:"iso_639_1"`
		Iso3166_1 string `json:"iso_3166_1"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		Site      string `json:"site"`
		Size      int    `json:"size"`
		Type      string `json:"type"`
	} `json:"results"`
}

// TVSeasonsVideosAppend type is a struct
// for videos in append to response.
type TVSeasonsVideosAppend struct {
	Videos struct {
		*TVSeasonsVideos
	} `json:"videos,omitempty"`
}

// GetTVSeasonsDetails get the TV season details by id.
//
// Supports append_to_response.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-details
func (c *Client) GetTVSeasonsDetails(id, s int, o map[string]string) (*TVSeasonsDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options,
	)
	t := TVSeasonsDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsChanges get the changes for a TV season.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-changes
func (c *Client) GetTVSeasonsChanges(id int, o map[string]string) (*TVSeasonsChanges, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf("%s%sseason/%d/changes?api_key=%s%s", baseURL, tvURL, id, c.APIKey, options)
	t := TVSeasonsChanges{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsCredits get the credits for TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-credits
func (c *Client) GetTVSeasonsCredits(id, s int, o map[string]string) (*TVSeasonsCredits, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/credits?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options,
	)
	t := TVSeasonsCredits{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsExternalIDs get the external ids for a TV season.
// We currently support the following external sources.
//
// Media Databases: TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-external-ids
func (c *Client) GetTVSeasonsExternalIDs(id, s int, o map[string]string) (*TVSeasonsExternalIDs, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/external_ids?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options,
	)
	t := TVSeasonsExternalIDs{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsImages get the images that belong to a TV season.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma seperated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-images
func (c *Client) GetTVSeasonsImages(id, s int, o map[string]string) (*TVSeasonsImages, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/images?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options,
	)
	t := TVSeasonsImages{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetTVSeasonsVideos get the videos that have been added to a TV season.
//
// https://developers.themoviedb.org/3/tv-seasons/get-tv-season-videos
func (c *Client) GetTVSeasonsVideos(id, s int, o map[string]string) (*TVSeasonsVideos, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d/videos?api_key=%s%s",
		baseURL, tvURL, id, tvSeasonURL, s, c.APIKey, options,
	)
	t := TVSeasonsVideos{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// TODO: Account States

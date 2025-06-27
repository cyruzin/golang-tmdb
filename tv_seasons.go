package tmdb

import "fmt"

// GetTVSeasonDetails get the TV season details by id.
//
// Supports append_to_response.
//
// https://developer.themoviedb.org/reference/tv-season-details
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

// GetTVSeasonChanges get the changes for a TV season.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developer.themoviedb.org/reference/tv-season-changes-by-id
func (c *Client) GetTVSeasonChanges(
	id int,
	urlOptions map[string]string,
) (*Changes, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sseason/%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvSeasonChanges := Changes{}
	if err := c.get(tmdbURL, &tvSeasonChanges); err != nil {
		return nil, err
	}
	return &tvSeasonChanges, nil
}

// GetTVSeasonCredits get the credits for TV season.
//
// https://developer.themoviedb.org/reference/tv-season-credits
func (c *Client) GetTVSeasonCredits(
	id int,
	seasonNumber int,
	urlOptions map[string]string,
) (*TVCredits, error) {
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
	tVSeasonCredits := TVCredits{}
	if err := c.get(tmdbURL, &tVSeasonCredits); err != nil {
		return nil, err
	}
	return &tVSeasonCredits, nil
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

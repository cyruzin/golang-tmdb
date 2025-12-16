package tmdb

import (
	"fmt"
)

// GetTVEpisodeDetails get the TV episode details by id.
//
// Supports append_to_response.
//
// https://developer.themoviedb.org/reference/tv-episode-details
func (c *Client) GetTVEpisodeDetails(
	id int,
	seasonNumber int,
	episodeNumber int,
	urlOptions map[string]string,
) (*TVEpisodeDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
		options,
	)
	tvEpisodeDetails := TVEpisodeDetails{}
	if err := c.get(tmdbURL, &tvEpisodeDetails); err != nil {
		return nil, err
	}
	return &tvEpisodeDetails, nil
}

// GetTVEpisodeChanges get the changes for a TV episode.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developer.themoviedb.org/reference/tv-episode-changes-by-id
func (c *Client) GetTVEpisodeChanges(
	id int,
	urlOptions map[string]string,
) (*TVEpisodeChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sepisode/%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvEpisodeChanges := TVEpisodeChanges{}
	if err := c.get(tmdbURL, &tvEpisodeChanges); err != nil {
		return nil, err
	}
	return &tvEpisodeChanges, nil
}

// GetTVEpisodeCredits get the credits (cast, crew and guest stars) for a TV episode.
//
// https://developer.themoviedb.org/reference/tv-episode-credits
func (c *Client) GetTVEpisodeCredits(
	id int,
	seasonNumber int,
	episodeNumber int,
) (*TVEpisodeCredits, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/credits?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeCredits := TVEpisodeCredits{}
	if err := c.get(tmdbURL, &tvEpisodeCredits); err != nil {
		return nil, err
	}
	return &tvEpisodeCredits, nil
}

// GetTVEpisodeExternalIDs get the external ids for a TV episode.
// We currently support the following external sources.
//
// Media Databases: IMDb ID, TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// *Defunct or no longer available as a service.
//
// https://developer.themoviedb.org/reference/tv-episode-external-ids
func (c *Client) GetTVEpisodeExternalIDs(
	id int,
	seasonNumber int,
	episodeNumber int,
) (*TVEpisodeExternalIDs, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/external_ids?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeExternalIDs := TVEpisodeExternalIDs{}
	if err := c.get(tmdbURL, &tvEpisodeExternalIDs); err != nil {
		return nil, err
	}
	return &tvEpisodeExternalIDs, nil
}

// GetTVEpisodeImages get the images that belong to a TV episode.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter.
// This should be a comma separated value like so: include_image_language=en,null.
//
// https://developer.themoviedb.org/reference/tv-episode-images
func (c *Client) GetTVEpisodeImages(
	id int,
	seasonNumber int,
	episodeNumber int,
) (*TVEpisodeImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/images?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeImages := TVEpisodeImages{}
	if err := c.get(tmdbURL, &tvEpisodeImages); err != nil {
		return nil, err
	}
	return &tvEpisodeImages, nil
}

// GetTVEpisodeTranslations get the translation data for an episode.
//
// https://developer.themoviedb.org/reference/tv-episode-translations
func (c *Client) GetTVEpisodeTranslations(
	id int,
	seasonNumber int,
	episodeNumber int,
) (*TVEpisodeTranslations, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/translations?api_key=%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
	)
	tvEpisodeTranslations := TVEpisodeTranslations{}
	if err := c.get(tmdbURL, &tvEpisodeTranslations); err != nil {
		return nil, err
	}
	return &tvEpisodeTranslations, nil
}

// GetTVEpisodeVideos get the videos that have been added to a TV episode.
//
// https://developer.themoviedb.org/reference/tv-episode-videos
func (c *Client) GetTVEpisodeVideos(
	id int,
	seasonNumber int,
	episodeNumber int,
	urlOptions map[string]string,
) (*VideoResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d%s%d%s%d/videos?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		tvSeasonURL,
		seasonNumber,
		tvEpisodeURL,
		episodeNumber,
		c.apiKey,
		options,
	)
	tvEpisodeVideos := VideoResults{}
	if err := c.get(tmdbURL, &tvEpisodeVideos); err != nil {
		return nil, err
	}
	return &tvEpisodeVideos, nil
}

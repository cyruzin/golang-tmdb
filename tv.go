package tmdb

import (
	"fmt"
	"net/http"
)

// GetTVDetails get the primary TV show details by id.
//
// Supports append_to_response.
//
// https://developer.themoviedb.org/reference/tv-series-details
func (c *Client) GetTVDetails(
	id int,
	urlOptions map[string]string,
) (*TVDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvDetails := TVDetails{}
	if err := c.get(tmdbURL, &tvDetails); err != nil {
		return nil, err
	}
	return &tvDetails, nil
}

// GetTVAccountStates grab the following account states for a session:
//
// TV show rating.
//
// If it belongs to your watchlist.
//
// If it belongs to your favourite list.
//
// https://developer.themoviedb.org/reference/tv-series-account-states
func (c *Client) GetTVAccountStates(
	id int,
	urlOptions map[string]string,
) (*TVAccountStates, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/account_states?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvAccountStates := TVAccountStates{}
	if err := c.get(tmdbURL, &tvAccountStates); err != nil {
		return nil, err
	}
	return &tvAccountStates, nil
}

// GetTVAggregateCredits get the aggregate credits (cast and crew) that have been added to a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-aggregate-credits
func (c *Client) GetTVAggregateCredits(
	id int,
	urlOptions map[string]string,
) (*TVAggregateCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/aggregate_credits?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvAggregateCredits := TVAggregateCredits{}
	if err := c.get(tmdbURL, &tvAggregateCredits); err != nil {
		return nil, err
	}
	return &tvAggregateCredits, nil
}

// GetTVAlternativeTitles get all of the alternative titles for a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-alternative-titles
func (c *Client) GetTVAlternativeTitles(
	id int,
	urlOptions map[string]string,
) (*TVAlternativeTitles, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_titles?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvAlternativeTitles := TVAlternativeTitles{}
	if err := c.get(tmdbURL, &tvAlternativeTitles); err != nil {
		return nil, err
	}
	return &tvAlternativeTitles, nil
}

// GetTVChanges get the changes for a TV show.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// TV show changes are different than movie changes in that there
// are some edits on seasons and episodes that will create a change
// entry at the show level. These can be found under the season
// and episode keys. These keys will contain a series_id and episode_id.
// You can use the and methods to look these up individually.
//
// https://developer.themoviedb.org/reference/tv-series-changes
func (c *Client) GetTVChanges(
	id int,
	urlOptions map[string]string,
) (*TVChanges, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tVChanges := TVChanges{}
	if err := c.get(tmdbURL, &tVChanges); err != nil {
		return nil, err
	}
	return &tVChanges, nil
}

// GetTVContentRatings get the list of content ratings (certifications) that have been added to a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-content-ratings
func (c *Client) GetTVContentRatings(
	id int,
	urlOptions map[string]string,
) (*TVContentRatings, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/content_ratings?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvContentRatings := TVContentRatings{}
	if err := c.get(tmdbURL, &tvContentRatings); err != nil {
		return nil, err
	}
	return &tvContentRatings, nil
}

// GetTVCredits get the credits (cast and crew) that have been added to a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-credits
func (c *Client) GetTVCredits(
	id int,
	urlOptions map[string]string,
) (*TVCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/credits?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvCredits := TVCredits{}
	if err := c.get(tmdbURL, &tvCredits); err != nil {
		return nil, err
	}
	return &tvCredits, nil
}

// GetTVEpisodeGroups get all of the episode groups that have been created for a TV show.
//
// With a group ID you can call the get TV episode group details method.
//
// https://developer.themoviedb.org/reference/tv-series-episode-groups
func (c *Client) GetTVEpisodeGroups(
	id int,
	urlOptions map[string]string,
) (*TVEpisodeGroups, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/episode_groups?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tVEpisodeGroups := TVEpisodeGroups{}
	if err := c.get(tmdbURL, &tVEpisodeGroups); err != nil {
		return nil, err
	}
	return &tVEpisodeGroups, nil
}

// GetTVExternalIDs get the external ids for a TV show.
//
// We currently support the following external sources.
//
// Media Databases: IMDb ID, TVDB ID, Freebase MID*, Freebase ID* TVRage ID*.
//
// Social IDs: Facebook, Instagram and Twitter.
//
// *Defunct or no longer available as a service.
//
// https://developer.themoviedb.org/reference/tv-series-external-ids
func (c *Client) GetTVExternalIDs(
	id int,
	urlOptions map[string]string,
) (*TVExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvExternalIDs := TVExternalIDs{}
	if err := c.get(tmdbURL, &tvExternalIDs); err != nil {
		return nil, err
	}
	return &tvExternalIDs, nil
}

// GetTVImages get the images that belong to a TV show.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter. This should be a comma
// separated value like so: include_image_language=en,null.
//
// https://developer.themoviedb.org/reference/tv-series-images
func (c *Client) GetTVImages(
	id int,
	urlOptions map[string]string,
) (*TVImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvImages := TVImages{}
	if err := c.get(tmdbURL, &tvImages); err != nil {
		return nil, err
	}
	return &tvImages, nil
}

// GetTVKeywords get the keywords that have been added to a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-keywords
func (c *Client) GetTVKeywords(
	id int,
) (*TVKeywords, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/keywords?api_key=%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
	)
	tvKeywords := TVKeywords{}
	if err := c.get(tmdbURL, &tvKeywords); err != nil {
		return nil, err
	}
	return &tvKeywords, nil
}

// GetTVRecommendations get the list of TV show recommendations for this item.
//
// https://developer.themoviedb.org/reference/tv-series-recommendations
func (c *Client) GetTVRecommendations(
	id int,
	urlOptions map[string]string,
) (*TVRecommendations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/recommendations?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvRecommendations := TVRecommendations{}
	if err := c.get(tmdbURL, &tvRecommendations); err != nil {
		return nil, err
	}
	return &tvRecommendations, nil
}

// GetTVReviews get the reviews for a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-reviews
func (c *Client) GetTVReviews(
	id int,
	urlOptions map[string]string,
) (*TVReviews, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/reviews?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvReviews := TVReviews{}
	if err := c.get(tmdbURL, &tvReviews); err != nil {
		return nil, err
	}
	return &tvReviews, nil
}

// GetTVScreenedTheatrically get a list of seasons or episodes that
// have been screened in a film festival or theatre.
//
// https://developer.themoviedb.org/reference/tv-series-screened-theatrically
func (c *Client) GetTVScreenedTheatrically(
	id int,
) (*TVScreenedTheatrically, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/screened_theatrically?api_key=%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
	)
	tvScreenedTheatrically := TVScreenedTheatrically{}
	if err := c.get(tmdbURL, &tvScreenedTheatrically); err != nil {
		return nil, err
	}
	return &tvScreenedTheatrically, nil
}

// GetTVSimilar a list of similar TV shows.
// These items are assembled by looking at keywords and genres.
//
// https://developer.themoviedb.org/reference/tv-series-similar
func (c *Client) GetTVSimilar(
	id int,
	urlOptions map[string]string,
) (*TVSimilar, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/similar?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tVSimilar := TVSimilar{}
	if err := c.get(tmdbURL, &tVSimilar); err != nil {
		return nil, err
	}
	return &tVSimilar, nil
}

// GetTVWatchProviders get a list of the availabilities per country by provider for a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-watch-providers
func (c *Client) GetTVWatchProviders(
	id int,
	urlOptions map[string]string,
) (*WatchProviderResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watch/providers?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvWatchProviders := WatchProviderResults{}
	if err := c.get(tmdbURL, &tvWatchProviders); err != nil {
		return nil, err
	}
	return &tvWatchProviders, nil
}

// GetTVTranslations get a list fo translations that have been created for a TV Show.
//
// https://developer.themoviedb.org/reference/tv-series-translations
func (c *Client) GetTVTranslations(
	id int,
	urlOptions map[string]string,
) (*TVTranslations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvTranslations := TVTranslations{}
	if err := c.get(tmdbURL, &tvTranslations); err != nil {
		return nil, err
	}
	return &tvTranslations, nil
}

// GetTVVideos get the videos that have been added to a TV show.
//
// https://developer.themoviedb.org/reference/tv-series-videos
func (c *Client) GetTVVideos(
	id int,
	urlOptions map[string]string,
) (*VideoResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/videos?api_key=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		options,
	)
	tvVideos := VideoResults{}
	if err := c.get(tmdbURL, &tvVideos); err != nil {
		return nil, err
	}
	return &tvVideos, nil
}

// GetTVLatest get the most newly created TV show.
//
// This is a live response and will continuously change.
//
// https://developer.themoviedb.org/reference/tv-series-latest-id
func (c *Client) GetTVLatest(
	urlOptions map[string]string,
) (*TVLatest, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvLatest := TVLatest{}
	if err := c.get(tmdbURL, &tvLatest); err != nil {
		return nil, err
	}
	return &tvLatest, nil
}

// GetTVAiringToday get a list of TV shows that are airing today.
// This query is purely day based as we do not currently support
// airing times.
//
// You can specify a  to offset the day calculation.
// Without a specified timezone, this query defaults
// to EST (Eastern Time UTC-05:00).
//
// https://developer.themoviedb.org/reference/tv-series-airing-today-list
func (c *Client) GetTVAiringToday(
	urlOptions map[string]string,
) (*TVAiringToday, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sairing_today?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvAiringToday := TVAiringToday{}
	if err := c.get(tmdbURL, &tvAiringToday); err != nil {
		return nil, err
	}
	return &tvAiringToday, nil
}

// GetTVOnTheAir get a list of shows that are currently on the air.
//
// This query looks for any TV show that has an episode with an
// air date in the next 7 days.
//
// https://developer.themoviedb.org/reference/tv-series-on-the-air-list
func (c *Client) GetTVOnTheAir(
	urlOptions map[string]string,
) (*TVOnTheAir, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%son_the_air?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvOnTheAir := TVOnTheAir{}
	if err := c.get(tmdbURL, &tvOnTheAir); err != nil {
		return nil, err
	}
	return &tvOnTheAir, nil
}

// GetTVPopular get a list of the current popular TV shows on TMDb.
// This list updates daily.
//
// https://developer.themoviedb.org/reference/tv-series-popular-list
func (c *Client) GetTVPopular(
	urlOptions map[string]string,
) (*TVPopular, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvPopular := TVPopular{}
	if err := c.get(tmdbURL, &tvPopular); err != nil {
		return nil, err
	}
	return &tvPopular, nil
}

// GetTVTopRated get a list of the top rated TV shows on TMDb.
//
// https://developer.themoviedb.org/reference/tv-series-top-rated-list
func (c *Client) GetTVTopRated(
	urlOptions map[string]string,
) (*TVTopRated, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stop_rated?api_key=%s%s",
		baseURL,
		tvURL,
		c.apiKey,
		options,
	)
	tvTopRated := TVTopRated{}
	if err := c.get(tmdbURL, &tvTopRated); err != nil {
		return nil, err
	}
	return &tvTopRated, nil
}

// PostTVShowRating rate a TV show.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developer.themoviedb.org/reference/authentication-how-do-i-generate-a-session-id
//
// https://developer.themoviedb.org/reference/tv-series-add-rating
func (c *Client) PostTVShowRating(
	id int,
	rating float32,
	urlOptions map[string]string,
) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	body := struct {
		Value float32 `json:"value"`
	}{Value: rating}
	tvShowRating := Response{}
	if err := c.request(
		tmdbURL,
		body,
		http.MethodPost,
		&tvShowRating,
	); err != nil {
		return nil, err
	}
	return &tvShowRating, nil
}

// DeleteTVShowRating remove your rating for a TV show.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developer.themoviedb.org/reference/authentication-how-do-i-generate-a-session-id
//
// https://developer.themoviedb.org/reference/tv-series-delete-rating
func (c *Client) DeleteTVShowRating(
	id int,
	urlOptions map[string]string,
) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		tvURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	tvShowRating := Response{}
	if err := c.request(
		tmdbURL,
		[]byte{},
		http.MethodDelete,
		&tvShowRating,
	); err != nil {
		return nil, err
	}
	return &tvShowRating, nil
}

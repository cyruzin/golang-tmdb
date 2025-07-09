package tmdb

import (
	"fmt"
	"net/http"
)

// GetMovieDetails get the primary information about a movie.
//
// https://developer.themoviedb.org/reference/movie-details
func (c *Client) GetMovieDetails(
	id int,
	urlOptions map[string]string,
) (*MovieDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieDetails := MovieDetails{}
	if err := c.get(tmdbURL, &movieDetails); err != nil {
		return nil, err
	}
	return &movieDetails, nil
}

// GetMovieAccountStates grab the following account states for a session:
//
// Movie rating.
//
// If it belongs to your watchlist.
//
// If it belongs to your favourite list.
//
// https://developer.themoviedb.org/reference/movie-account-states
func (c *Client) GetMovieAccountStates(
	id int,
	urlOptions map[string]string,
) (*MovieAccountStates, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/account_states?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieAccountStates := MovieAccountStates{}
	if err := c.get(tmdbURL, &movieAccountStates); err != nil {
		return nil, err
	}
	return &movieAccountStates, nil
}

// GetMovieAlternativeTitles get all of the alternative titles for a movie.
//
// https://developer.themoviedb.org/reference/movie-alternative-titles
func (c *Client) GetMovieAlternativeTitles(
	id int,
	urlOptions map[string]string,
) (*MovieAlternativeTitles, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_titles?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieAlternativeTitles := MovieAlternativeTitles{}
	if err := c.get(tmdbURL, &movieAlternativeTitles); err != nil {
		return nil, err
	}
	return &movieAlternativeTitles, nil
}

// GetMovieChanges get the changes for a movie.
//
// By default only the last 24 hours are returned.
// You can query up to 14 days in a single query by using
// the start_date and end_date query parameters.
//
// https://developer.themoviedb.org/reference/movie-changes
func (c *Client) GetMovieChanges(
	id int,
	urlOptions map[string]string,
) (*Changes, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/changes?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieChanges := Changes{}
	if err := c.get(tmdbURL, &movieChanges); err != nil {
		return nil, err
	}
	return &movieChanges, nil
}

// GetMovieCredits get the cast and crew for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-credits
func (c *Client) GetMovieCredits(
	id int,
	urlOptions map[string]string,
) (*MovieCredits, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf("%s%s%d/credits?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieCredits := MovieCredits{}
	if err := c.get(tmdbURL, &movieCredits); err != nil {
		return nil, err
	}
	return &movieCredits, nil
}

// GetMovieExternalIDs get the external ids for a movie.
//
// We currently support the following external sources.
//
// Media Databases: IMDb ID.
//
// Social IDs: Facebook, Instagram and Twitter.
//
// https://developers.themoviedb.org/3/movies/get-movie-external-ids
func (c *Client) GetMovieExternalIDs(
	id int,
	urlOptions map[string]string,
) (*MovieExternalIDs, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/external_ids?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieExternalIDs := MovieExternalIDs{}
	if err := c.get(tmdbURL, &movieExternalIDs); err != nil {
		return nil, err
	}
	return &movieExternalIDs, nil
}

// GetMovieImages get the images that belong to a movie.
//
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (especially useful for backdrops)
// you can use the include_image_language parameter. This should be a comma
// separated value like so: include_image_language=en,null.
//
// https://developers.themoviedb.org/3/movies/get-movie-images
func (c *Client) GetMovieImages(
	id int,
	urlOptions map[string]string,
) (*MovieImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieImages := MovieImages{}
	if err := c.get(tmdbURL, &movieImages); err != nil {
		return nil, err
	}
	return &movieImages, nil
}

// GetMovieKeywords get the keywords that have been added to a movie.
//
// https://developer.themoviedb.org/reference/movie-keywords
func (c *Client) GetMovieKeywords(id int) (*IDKeywordsList, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/keywords?api_key=%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
	)
	movieKeywords := IDKeywordsList{}
	if err := c.get(tmdbURL, &movieKeywords); err != nil {
		return nil, err
	}
	return &movieKeywords, nil
}

// GetMovieReleaseDates get the release date along with the certification for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-release-dates
func (c *Client) GetMovieReleaseDates(
	id int,
) (*IDMovieReleaseDateResults, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/release_dates?api_key=%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
	)
	movieReleaseDates := IDMovieReleaseDateResults{}
	if err := c.get(tmdbURL, &movieReleaseDates); err != nil {
		return nil, err
	}
	return &movieReleaseDates, nil
}

// GetMovieVideos get the videos that have been added to a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-videos
func (c *Client) GetMovieVideos(
	id int,
	urlOptions map[string]string,
) (*VideoResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/videos?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieVideos := VideoResults{}
	if err := c.get(tmdbURL, &movieVideos); err != nil {
		return nil, err
	}
	return &movieVideos, nil
}

// GetMovieWatchProviders get a list of the availabilities per country by provider for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-watch-providers
func (c *Client) GetMovieWatchProviders(
	id int,
	urlOptions map[string]string,
) (*IDWatchProviderResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/watch/providers?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieWatchProviders := IDWatchProviderResults{}
	if err := c.get(tmdbURL, &movieWatchProviders); err != nil {
		return nil, err
	}
	return &movieWatchProviders, nil
}

// GetMovieTranslations get a list of translations that have been created for a movie.
//
// https://developers.themoviedb.org/3/movies/get-movie-translations
func (c *Client) GetMovieTranslations(
	id int,
	urlOptions map[string]string,
) (*MovieTranslations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieTranslations := MovieTranslations{}
	if err := c.get(tmdbURL, &movieTranslations); err != nil {
		return nil, err
	}
	return &movieTranslations, nil
}

// GetMovieRecommendations get a list of recommended movies for a movie.
//
// https://developer.themoviedb.org/reference/movie-recommendations
func (c *Client) GetMovieRecommendations(
	id int,
	urlOptions map[string]string,
) (*PaginatedMovieMediaResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/recommendations?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieRecommendations := PaginatedMovieMediaResults{}
	if err := c.get(tmdbURL, &movieRecommendations); err != nil {
		return nil, err
	}
	return &movieRecommendations, nil
}

// GetMovieSimilar get a list of similar movies.
//
// This is not the same as the "Recommendation" system you see on the website.
// These items are assembled by looking at keywords and genres.
//
// https://developers.themoviedb.org/3/movies/get-similar-movies
func (c *Client) GetMovieSimilar(
	id int,
	urlOptions map[string]string,
) (*PaginatedMovieResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/similar?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieSimilar := PaginatedMovieResults{}
	if err := c.get(tmdbURL, &movieSimilar); err != nil {
		return nil, err
	}
	return &movieSimilar, nil
}

// GetMovieReviews get the user reviews for a movie.
//
// https://developer.themoviedb.org/reference/movie-reviews
func (c *Client) GetMovieReviews(
	id int,
	urlOptions map[string]string,
) (*IDPaginatedReviewsResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/reviews?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieReviews := IDPaginatedReviewsResults{}
	if err := c.get(tmdbURL, &movieReviews); err != nil {
		return nil, err
	}
	return &movieReviews, nil
}

// GetMovieLists get a list of lists that this movie belongs to.
//
// https://developer.themoviedb.org/reference/movie-lists
func (c *Client) GetMovieLists(
	id int,
	urlOptions map[string]string,
) (*IDPaginatedMovieListResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/lists?api_key=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		options,
	)
	movieLists := IDPaginatedMovieListResults{}
	if err := c.get(tmdbURL, &movieLists); err != nil {
		return nil, err
	}
	return &movieLists, nil
}

// GetMovieLatest get the most newly created movie.
//
// This is a live response and will continuously change.
//
// https://developers.themoviedb.org/3/movies/get-latest-movie
func (c *Client) GetMovieLatest(
	urlOptions map[string]string,
) (*MovieLatest, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%slatest?api_key=%s%s",
		baseURL,
		movieURL,
		c.apiKey,
		options,
	)
	movieLastest := MovieLatest{}
	if err := c.get(tmdbURL, &movieLastest); err != nil {
		return nil, err
	}
	return &movieLastest, nil
}

// PostMovieRating rate a movie.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
//
// https://developers.themoviedb.org/3/movies/rate-movie
func (c *Client) PostMovieRating(
	id int,
	rating float32,
	urlOptions map[string]string,
) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	body := struct {
		Value float32 `json:"value"`
	}{Value: rating}
	response := Response{}
	if err := c.request(
		tmdbURL,
		body,
		http.MethodPost,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteMovieRating remove your rating for a movie.
//
// A valid session or guest session ID is required.
//
// You can read more about how this works:
// https://developers.themoviedb.org/3/authentication/how-do-i-generate-a-session-id
//
// https://developers.themoviedb.org/3/movies/delete-movie-rating
func (c *Client) DeleteMovieRating(
	id int,
	urlOptions map[string]string,
) (*Response, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/rating?api_key=%s&session_id=%s%s",
		baseURL,
		movieURL,
		id,
		c.apiKey,
		c.sessionID,
		options,
	)
	response := Response{}
	if err := c.request(
		tmdbURL,
		[]byte{},
		http.MethodDelete,
		&response,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

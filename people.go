package tmdb

import "fmt"

// GetPersonDetails get the primary person details by id.
//
// Supports append_to_response.
//
// https://developer.themoviedb.org/reference/person-details
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

// GetPersonChanges get the changes for a person.
// By default only the last 24 hours are returned.
//
// You can query up to 14 days in a single query by
// using the start_date and end_date query parameters.
//
// https://developer.themoviedb.org/reference/person-changes
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

// GetPersonMovieCredits get the movie credits for a person.
//
// https://developer.themoviedb.org/reference/person-movie-credits
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

// GetPersonTVCredits get the TV show credits for a person.
//
// https://developer.themoviedb.org/reference/person-tv-credits
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

// GetPersonCombinedCredits get the movie and TV credits together in a single response.
//
// https://developer.themoviedb.org/reference/person-combined-credits
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

// GetPersonExternalIDs get the external ids for a person.
// We currently support the following external sources.
//
// External Sources: IMDb ID, Facebook, Freebase MID, Freebase ID,
// Instagram, TVRage ID, Twitter.
//
// https://developer.themoviedb.org/reference/person-external-ids
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

// GetPersonLatest get the most newly created person.
// This is a live response and will continuously change.
//
// https://developer.themoviedb.org/reference/person-latest-id
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
// https://developer.themoviedb.org/reference/person-popular-list
func (c *Client) GetPersonPopular(
	urlOptions map[string]string,
) (*PersonPopular, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%spopular?api_key=%s%s",
		baseURL,
		personURL,
		c.apiKey,
		options,
	)
	personPopular := PersonPopular{}
	if err := c.get(tmdbURL, &personPopular); err != nil {
		return nil, err
	}
	return &personPopular, nil
}

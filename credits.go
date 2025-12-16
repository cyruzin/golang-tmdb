package tmdb

import (
	"fmt"
)

// GetCreditDetails get a movie or TV credit details by id.
//
// https://developer.themoviedb.org/reference/credit-details
func (c *Client) GetCreditDetails(
	id string,
) (*CreditsDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%s?api_key=%s",
		baseURL,
		creditURL,
		id,
		c.apiKey,
	)
	creditsDetails := CreditsDetails{}
	if err := c.get(tmdbURL, &creditsDetails); err != nil {
		return nil, err
	}
	return &creditsDetails, nil
}

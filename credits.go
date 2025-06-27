package tmdb

import (
	"fmt"
)

// CreditsDetails type is a struct for credits JSON response.
type CreditsDetails struct {
	CreditType string          `json:"credit_type"`
	Department string          `json:"department"`
	Job        string          `json:"job"`
	Media      CreditMedia     `json:"media"`
	MediaType  MediaType       `json:"media_type"`
	ID         string          `json:"id"`
	Person     PersonMediaBase `json:"person"`
}

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

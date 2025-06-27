package tmdb

import (
	"fmt"
)

// CreditsDetails type is a struct for credits JSON response.
type CreditsDetails struct {
	CreditType string      `json:"credit_type"`
	Department string      `json:"department"`
	Job        string      `json:"job"`
	Media      CreditMedia `json:"media"`
	MediaType  MediaType   `json:"media_type"`
	ID         string      `json:"id"`
	Person     PersonMedia `json:"person"`
}

type CreditSeason struct {
	Season
	MediaType MediaType `json:"media_type"`
	ShowID    int64     `json:"show_id"`
}

// GetCreditDetails get a movie or TV credit details by id.
//
// https://developers.themoviedb.org/3/credits/get-credit-details
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

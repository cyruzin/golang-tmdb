package tmdb

import "fmt"

// GetReviewDetails get review details by id.
//
// https://developer.themoviedb.org/reference/review-details
func (c *Client) GetReviewDetails(
	id string,
) (*ReviewDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s/review/%s?api_key=%s",
		baseURL,
		id,
		c.apiKey,
	)
	reviewDetails := ReviewDetails{}
	if err := c.get(tmdbURL, &reviewDetails); err != nil {
		return nil, err
	}
	return &reviewDetails, nil
}

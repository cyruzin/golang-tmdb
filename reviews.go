package tmdb

import "fmt"

// ReviewDetails type is a struct for details JSON response.
type ReviewDetails struct {
	ID            string `json:"id"`
	Author        string `json:"author"`
	AuthorDetails struct {
		AvatarPath string  `json:"avatar_path"`
		Name       string  `json:"name"`
		Rating     float32 `json:"rating"`
		Username   string  `json:"username"`
	} `json:"author_details"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Iso639_1   string `json:"iso_639_1"`
	MediaID    int64  `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
	URL        string `json:"url"`
}

// GetReviewDetails get review details by id.
//
// https://developers.themoviedb.org/3/reviews/get-review-details
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

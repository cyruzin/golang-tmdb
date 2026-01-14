package tmdb

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
func (c *Client) GetReviewDetails(id string) (*ReviewDetails, error) {
	url := c.baseURL.JoinPath("review", id)
	var reviewDetails ReviewDetails
	if err := c.get(url, nil, &reviewDetails); err != nil {
		return nil, err
	}
	return &reviewDetails, nil
}

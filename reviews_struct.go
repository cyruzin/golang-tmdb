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

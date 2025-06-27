package tmdb

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

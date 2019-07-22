package tmdb

// AccountDetails type is a struct for details JSON response.
type AccountDetails struct {
	Avatar struct {
		Gravatar struct {
			Hash string `json:"hash"`
		} `json:"gravatar"`
	} `json:"avatar"`
	ID           int64  `json:"id"`
	Iso639_1     string `json:"iso_639_1"`
	Iso3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

// GetAccountDetails get your account details.
//
// https://developers.themoviedb.org/3/account/get-account-details
// func (c *Client) GetAccountDetails(o map[string]string) (*AccountDetails, error) {
// 	options := c.fmtOptions(o)
// 	tmdbURL := fmt.Sprintf(
// 		"%s/account?api_key=%s%s", baseURL, c.apiKey, options,
// 	)
// 	fmt.Println(tmdbURL)
// 	a := AccountDetails{}
// 	err := c.get(tmdbURL, &a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

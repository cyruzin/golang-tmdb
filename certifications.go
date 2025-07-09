package tmdb

import "fmt"

// GetCertificationMovie get an up to date list of the
// officially supported movie certifications on TMDb.
//
// https://developer.themoviedb.org/reference/certification-movie-list
func (c *Client) GetCertificationMovie() (
	*Certifications,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL,
		movieURL,
		c.apiKey,
	)
	certificationMovie := Certifications{}
	if err := c.get(tmdbURL, &certificationMovie); err != nil {
		return nil, err
	}
	return &certificationMovie, nil
}

// GetCertificationTV get an up to date list of the
// officially supported TV show certifications on TMDb.
//
// https://developer.themoviedb.org/reference/certifications-tv-list
func (c *Client) GetCertificationTV() (
	*Certifications,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL,
		tvURL,
		c.apiKey,
	)
	certificationTV := Certifications{}
	if err := c.get(tmdbURL, &certificationTV); err != nil {
		return nil, err
	}
	return &certificationTV, nil
}

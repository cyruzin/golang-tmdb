package tmdb

// Certification type is a struct for a single certification JSON response.
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int    `json:"order"`
}

// Certifications type is a struct for movie and tv certifications JSON response.
type Certifications struct {
	Certifications map[string][]Certification `json:"certifications"`
}

// GetCertificationMovie get an up to date list of the
// officially supported movie certifications on TMDb.
//
// https://developer.themoviedb.org/reference/certification-movie-list
func (c *Client) GetCertificationMovie() (*Certifications, error) {
	url := c.baseURL.JoinPath("certification", moviePath, "list")
	var certificationMovie Certifications
	if err := c.get(url, nil, &certificationMovie); err != nil {
		return nil, err
	}
	return &certificationMovie, nil
}

// GetCertificationTV get an up to date list of the
// officially supported TV show certifications on TMDb.
//
// https://developer.themoviedb.org/reference/certifications-tv-list
func (c *Client) GetCertificationTV() (*Certifications, error) {
	url := c.baseURL.JoinPath("certification", tvPath, "list")
	var certificationTV Certifications
	if err := c.get(url, nil, &certificationTV); err != nil {
		return nil, err
	}
	return &certificationTV, nil
}

package tmdb

import "fmt"

// Certification type is a struct for a single certification JSON response.
type Certification struct {
	Certification string `json:"certification"`
	Meaning       string `json:"meaning"`
	Order         int    `json:"order"`
}

// CertificationMovie type is a struct for movie certifications JSON response.
type CertificationMovie struct {
	Certifications struct {
		US []struct {
			*Certification
		} `json:"US"`
		CA []struct {
			*Certification
		} `json:"CA"`
		AU []struct {
			*Certification
		} `json:"AU"`
		DE []struct {
			*Certification
		} `json:"DE"`
		FR []struct {
			*Certification
		} `json:"FR"`
		NZ []struct {
			*Certification
		} `json:"NZ"`
		IN []struct {
			*Certification
		} `json:"IN"`
		GB []struct {
			*Certification
		} `json:"GB"`
		NL []struct {
			*Certification
		} `json:"NL"`
		BR []struct {
			*Certification
		} `json:"BR"`
		FI []struct {
			*Certification
		} `json:"FI"`
		BG []struct {
			*Certification
		} `json:"BG"`
		ES []struct {
			*Certification
		} `json:"ES"`
		PH []struct {
			*Certification
		} `json:"PH"`
		PT []struct {
			*Certification
		} `json:"PT"`
	} `json:"certifications"`
}

// CertificationTV type is a struct for tv certifications JSON response.
type CertificationTV struct {
	Certifications struct {
		RU []struct {
			*Certification
		} `json:"RU"`
		US []struct {
			*Certification
		} `json:"US"`
		CA []struct {
			*Certification
		} `json:"CA"`
		AU []struct {
			*Certification
		} `json:"AU"`
		FR []struct {
			*Certification
		} `json:"FR"`
		DE []struct {
			*Certification
		} `json:"DE"`
		TH []struct {
			*Certification
		} `json:"TH"`
		KR []struct {
			*Certification
		} `json:"KR"`
		GB []struct {
			*Certification
		} `json:"GB"`
		BR []struct {
			*Certification
		} `json:"BR"`
	} `json:"certifications"`
}

// GetCertificationMovie get an up to date list of the
// officially supported movie certifications on TMDb.
//
// https://developers.themoviedb.org/3/certifications/get-movie-certifications
func (c *Client) GetCertificationMovie() (
	*CertificationMovie, error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL, movieURL, c.apiKey,
	)
	m := CertificationMovie{}
	err := c.get(tmdbURL, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// GetCertificationTV get an up to date list of the
// officially supported TV show certifications on TMDb.
//
// https://developers.themoviedb.org/3/certifications/get-tv-certifications
func (c *Client) GetCertificationTV() (
	*CertificationTV, error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL, tvURL, c.apiKey,
	)
	t := CertificationTV{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

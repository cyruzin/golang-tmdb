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
		AU []struct {
			*Certification
		} `json:"AU"`
		BG []struct {
			*Certification
		} `json:"BG"`
		BR []struct {
			*Certification
		} `json:"BR"`
		CA []struct {
			*Certification
		} `json:"CA"`
		CA_QC []struct {
			*Certification
		} `json:"CA-QC"`
		DE []struct {
			*Certification
		} `json:"DE"`
		DK []struct {
			*Certification
		} `json:"DK"`
		ES []struct {
			*Certification
		} `json:"ES"`
		FI []struct {
			*Certification
		} `json:"FI"`
		FR []struct {
			*Certification
		} `json:"FR"`
		GB []struct {
			*Certification
		} `json:"GB"`
		HU []struct {
			*Certification
		} `json:"HU"`
		IN []struct {
			*Certification
		} `json:"IN"`
		IT []struct {
			*Certification
		} `json:"IT"`
		LT []struct {
			*Certification
		} `json:"LT"`
		MY []struct {
			*Certification
		} `json:"MY"`
		NL []struct {
			*Certification
		} `json:"NL"`
		NO []struct {
			*Certification
		} `json:"NO"`
		NZ []struct {
			*Certification
		} `json:"NZ"`
		PH []struct {
			*Certification
		} `json:"PH"`
		PT []struct {
			*Certification
		} `json:"PT"`
		RU []struct {
			*Certification
		} `json:"RU"`
		SE []struct {
			*Certification
		} `json:"SE"`
		US []struct {
			*Certification
		} `json:"US"`
	} `json:"certifications"`
}

// GetCertificationMovie get an up to date list of the
// officially supported movie certifications on TMDb.
//
// https://developers.themoviedb.org/3/certifications/get-movie-certifications
func (c *Client) GetCertificationMovie() (
	*CertificationMovie,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL,
		movieURL,
		c.apiKey,
	)
	certificationMovie := CertificationMovie{}
	if err := c.get(tmdbURL, &certificationMovie); err != nil {
		return nil, err
	}
	return &certificationMovie, nil
}

// CertificationTV type is a struct for tv certifications JSON response.
type CertificationTV struct {
	Certifications struct {
		AU []struct {
			*Certification
		} `json:"AU"`
		BR []struct {
			*Certification
		} `json:"BR"`
		CA []struct {
			*Certification
		} `json:"CA"`
		CA_QC []struct {
			*Certification
		} `json:"CA-QC"`
		DE []struct {
			*Certification
		} `json:"DE"`
		ES []struct {
			*Certification
		} `json:"ES"`
		FR []struct {
			*Certification
		} `json:"FR"`
		GB []struct {
			*Certification
		} `json:"GB"`
		HU []struct {
			*Certification
		} `json:"HU"`
		KR []struct {
			*Certification
		} `json:"KR"`
		LT []struct {
			*Certification
		} `json:"LT"`
		NL []struct {
			*Certification
		} `json:"NL"`
		PH []struct {
			*Certification
		} `json:"PH"`
		PT []struct {
			*Certification
		} `json:"PT"`
		RU []struct {
			*Certification
		} `json:"RU"`
		SK []struct {
			*Certification
		} `json:"SK"`
		TH []struct {
			*Certification
		} `json:"TH"`
		US []struct {
			*Certification
		} `json:"US"`
	} `json:"certifications"`
}

// GetCertificationTV get an up to date list of the
// officially supported TV show certifications on TMDb.
//
// https://developers.themoviedb.org/3/certifications/get-tv-certifications
func (c *Client) GetCertificationTV() (
	*CertificationTV,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s/certification%slist?api_key=%s",
		baseURL,
		tvURL,
		c.apiKey,
	)
	certificationTV := CertificationTV{}
	if err := c.get(tmdbURL, &certificationTV); err != nil {
		return nil, err
	}
	return &certificationTV, nil
}

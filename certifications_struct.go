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

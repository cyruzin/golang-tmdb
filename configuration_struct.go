package tmdb

// ConfigurationAPI type is a struct for api configuration JSON response.
type ConfigurationAPI struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	} `json:"images"`
	ChangeKeys []string `json:"change_keys"`
}

// ConfigurationCountries type is a struct for countries configuration JSON response.
type ConfigurationCountries []struct {
	Iso3166_1   string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
	NativeName  string `json:"native_name"`
}

// ConfigurationJobs type is a struct for jobs configuration JSON response.
type ConfigurationJobs []struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// ConfigurationLanguages type is a struct for languages configuration JSON response.
type ConfigurationLanguages []struct {
	Iso639_1    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

// ConfigurationPrimaryTranslations type is a struct for
// primary translations configuration JSON response.
type ConfigurationPrimaryTranslations []string

// ConfigurationTimezones type is a struct for timezones
// configuration JSON response.
type ConfigurationTimezones []struct {
	Iso3166_1 string   `json:"iso_3166_1"`
	Zones     []string `json:"zones"`
}

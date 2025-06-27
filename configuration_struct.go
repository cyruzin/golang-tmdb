package tmdb

type ImagesConfiguration struct {
	BaseURL       string   `json:"base_url"`
	SecureBaseURL string   `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}

// ConfigurationAPI type is a struct for api configuration JSON response.
type ConfigurationAPI struct {
	Images     ImagesConfiguration `json:"images"`
	ChangeKeys []string            `json:"change_keys"`
}

// ConfigurationCountries type is a struct for countries configuration JSON response.
type ConfigurationCountries []Country

// ConfigurationJobs type is a struct for jobs configuration JSON response.
type ConfigurationJobs []Jobs

// ConfigurationLanguages type is a struct for languages configuration JSON response.
type ConfigurationLanguages []SpokenLanguage

// ConfigurationPrimaryTranslations type is a struct for
// primary translations configuration JSON response.
type ConfigurationPrimaryTranslations []string

type Timezones struct {
	Iso3166_1 string   `json:"iso_3166_1"`
	Zones     []string `json:"zones"`
}

// ConfigurationTimezones type is a struct for timezones
// configuration JSON response.
type ConfigurationTimezones []Timezones

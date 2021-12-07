package tmdb

import "fmt"

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

// GetConfigurationAPI get the system wide configuration information.
//
// Some elements of the API require some knowledge of
// this configuration data. The purpose of this is to
// try and keep the actual API responses as light as possible.
// It is recommended you cache this data within your application
// and check for updates every few days.
//
// This method currently holds the data relevant to building image
// URLs as well as the change key map.
//
// To build an image URL, you will need 3 pieces of data. The base_url,
// size and file_path. Simply combine them all and you will have a fully
// qualified URL. Here’s an example URL:
//
// https://image.tmdb.org/t/p/w500/8uO0gUM8aNqYLs1OsTBQiXu0fEv.jpg
//
// The configuration method also contains the list of change keys which
// can be useful if you are building an app that consumes data from the
// change feed.
//
// https://developers.themoviedb.org/3/configuration/get-api-configuration
func (c *Client) GetConfigurationAPI() (*ConfigurationAPI, error) {
	tmdbURL := fmt.Sprintf(
		"%s/configuration?api_key=%s",
		baseURL,
		c.apiKey,
	)
	configurationAPI := ConfigurationAPI{}
	if err := c.get(tmdbURL, &configurationAPI); err != nil {
		return nil, err
	}
	return &configurationAPI, nil
}

// ConfigurationCountries type is a struct for countries configuration JSON response.
type ConfigurationCountries []struct {
	Iso3166_1   string `json:"iso_3166_1"`
	EnglishName string `json:"english_name"`
	NativeName  string `json:"native_name"`
}

// GetConfigurationCountries get the list of countries
// (ISO 3166-1 tags) used throughout TMDb.
//
// https://developers.themoviedb.org/3/configuration/get-countries
func (c *Client) GetConfigurationCountries() (
	*ConfigurationCountries,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s%scountries?api_key=%s",
		baseURL,
		configurationURL,
		c.apiKey,
	)
	configurationCountries := ConfigurationCountries{}
	if err := c.get(tmdbURL, &configurationCountries); err != nil {
		return nil, err
	}
	return &configurationCountries, nil
}

// ConfigurationJobs type is a struct for jobs configuration JSON response.
type ConfigurationJobs []struct {
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// GetConfigurationJobs get a list of the jobs and departments we use on TMDb.
//
// https://developers.themoviedb.org/3/configuration/get-jobs
func (c *Client) GetConfigurationJobs() (*ConfigurationJobs, error) {
	tmdbURL := fmt.Sprintf(
		"%s%sjobs?api_key=%s",
		baseURL,
		configurationURL,
		c.apiKey,
	)
	configurationJobs := ConfigurationJobs{}
	if err := c.get(tmdbURL, &configurationJobs); err != nil {
		return nil, err
	}
	return &configurationJobs, nil
}

// ConfigurationLanguages type is a struct for languages configuration JSON response.
type ConfigurationLanguages []struct {
	Iso639_1    string `json:"iso_639_1"`
	EnglishName string `json:"english_name"`
	Name        string `json:"name"`
}

// GetConfigurationLanguages get the list of languages
// (ISO 639-1 tags) used throughout TMDb.
//
// https://developers.themoviedb.org/3/configuration/get-languages
func (c *Client) GetConfigurationLanguages() (
	*ConfigurationLanguages,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s%slanguages?api_key=%s",
		baseURL,
		configurationURL,
		c.apiKey,
	)
	configurationLanguages := ConfigurationLanguages{}
	if err := c.get(tmdbURL, &configurationLanguages); err != nil {
		return nil, err
	}
	return &configurationLanguages, nil
}

// ConfigurationPrimaryTranslations type is a struct for
// primary translations configuration JSON response.
type ConfigurationPrimaryTranslations []string

// GetConfigurationPrimaryTranslations get a list of the officially
// supported translations on TMDb.
//
// While it's technically possible to add a translation in any one
// of the languages we have added to TMDb (we don't restrict content),
// the ones listed in this method are the ones we also support for
// localizing the website with which means they are what we refer to
// as the "primary" translations.
//
// These are all specified as IETF tags to identify the languages
// we use on TMDb. There is one exception which is image languages.
// They are currently only designated by a ISO-639-1 tag. This is
// a planned upgrade for the future.
//
// We're always open to adding more if you think one should be added.
// You can ask about getting a new primary translation added by posting
// on the forums.
//
// One more thing to mention, these are the translations that map to
// our website translation project.
//
// https://developers.themoviedb.org/3/configuration/get-primary-translations
func (c *Client) GetConfigurationPrimaryTranslations() (
	*ConfigurationPrimaryTranslations,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s%sprimary_translations?api_key=%s",
		baseURL, configurationURL, c.apiKey,
	)
	configurationPrimaryTranslations := ConfigurationPrimaryTranslations{}
	if err := c.get(tmdbURL, &configurationPrimaryTranslations); err != nil {
		return nil, err
	}
	return &configurationPrimaryTranslations, nil
}

// ConfigurationTimezones type is a struct for timezones
// configuration JSON response.
type ConfigurationTimezones []struct {
	Iso3166_1 string   `json:"iso_3166_1"`
	Zones     []string `json:"zones"`
}

// GetConfigurationTimezones get the list of timezones
// used throughout TMDb.
//
// https://developers.themoviedb.org/3/configuration/get-timezones
func (c *Client) GetConfigurationTimezones() (
	*ConfigurationTimezones,
	error,
) {
	tmdbURL := fmt.Sprintf(
		"%s%stimezones?api_key=%s",
		baseURL,
		configurationURL,
		c.apiKey,
	)
	configurationTimeZones := ConfigurationTimezones{}
	if err := c.get(tmdbURL, &configurationTimeZones); err != nil {
		return nil, err
	}
	return &configurationTimeZones, nil
}

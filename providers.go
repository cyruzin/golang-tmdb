package tmdb

// WatchRegionList type is a struct for watch region list JSON response.
type WatchRegionList struct {
	Regions []struct {
		Iso3166_1   string `json:"iso_3166_1"`
		EnglishName string `json:"english_name"`
		NativeName  string `json:"native_name"`
	} `json:"results"`
}

// WatchProviderList type is a struct for watch provider list JSON response.
type WatchProviderList struct {
	Providers []struct {
		ID                int            `json:"id"`
		Name              string         `json:"name"`
		DisplayPriorities map[string]int `json:"display_priorities"`
		DisplayPriority   int            `json:"display_priority"`
		LogoPath          string         `json:"logo_path"`
		ProviderName      string         `json:"provider_name"`
		ProviderID        int            `json:"provider_id"`
	} `json:"results"`
}

// GetAvailableWatchProviderRegions get a list of all of the countries we have watch provider (OTT/streaming) data for.
//
// https://developers.themoviedb.org/3/watch-providers/get-available-regions
func (c *Client) GetAvailableWatchProviderRegions(urlOptions map[string]string) (*WatchRegionList, error) {
	url := c.baseURL.JoinPath(watchProvidersPath, "regions")
	var watchRegionList WatchRegionList
	if err := c.get(url, urlOptions, &watchRegionList); err != nil {
		return nil, err
	}
	return &watchRegionList, nil
}

// GetWatchProvidersMovie get a list of the watch provider (OTT/streaming) data we have available for movies.
// You can specify a watch_region param if you want to further filter the list by country.
//
// https://developers.themoviedb.org/3/watch-providers/get-movie-providers
func (c *Client) GetWatchProvidersMovie(urlOptions map[string]string) (*WatchProviderList, error) {
	url := c.baseURL.JoinPath(watchProvidersPath, "movie")
	var watchProvider WatchProviderList
	if err := c.get(url, urlOptions, &watchProvider); err != nil {
		return nil, err
	}
	return &watchProvider, nil
}

// GetWatchProvidersTv get a list of the watch provider (OTT/streaming) data we have available for TV series.
// You can specify a watch_region param if you want to further filter the list by country.
//
// https://developers.themoviedb.org/3/watch-providers/get-tv-providers
func (c *Client) GetWatchProvidersTv(urlOptions map[string]string) (*WatchProviderList, error) {
	url := c.baseURL.JoinPath(watchProvidersPath, "tv")
	var watchProvider WatchProviderList
	if err := c.get(url, urlOptions, &watchProvider); err != nil {
		return nil, err
	}
	return &watchProvider, nil
}

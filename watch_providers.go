package tmdb

import "fmt"

// WatchRegionList type is a struct for watch region list JSON response.

// GetAvailableWatchProviderRegions get a list of all of the countries we have watch provider (OTT/streaming) data for.
//
// https://developer.themoviedb.org/reference/watch-providers-available-regions
func (c *Client) GetAvailableWatchProviderRegions(
	urlOptions map[string]string,
) (*CountryWatchProviderResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%sregions?api_key=%s%s",
		baseURL,
		watchProvidersURL,
		c.apiKey,
		options,
	)
	watchRegionList := CountryWatchProviderResults{}
	if err := c.get(tmdbURL, &watchRegionList); err != nil {
		return nil, err
	}
	return &watchRegionList, nil
}

// GetWatchProvidersMovie get a list of the watch provider (OTT/streaming) data we have available for movies.
// You can specify a watch_region param if you want to further filter the list by country.
//
// https://developer.themoviedb.org/reference/watch-providers-movie-list
func (c *Client) GetWatchProvidersMovie(
	urlOptions map[string]string,
) (*MediaWatchProviderResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%smovie?api_key=%s%s",
		baseURL,
		watchProvidersURL,
		c.apiKey,
		options,
	)
	watchProvider := MediaWatchProviderResults{}
	if err := c.get(tmdbURL, &watchProvider); err != nil {
		return nil, err
	}
	return &watchProvider, nil
}

// GetWatchProvidersTv get a list of the watch provider (OTT/streaming) data we have available for TV series.
// You can specify a watch_region param if you want to further filter the list by country.
//
// https://developer.themoviedb.org/reference/watch-provider-tv-list
func (c *Client) GetWatchProvidersTv(
	urlOptions map[string]string,
) (*MediaWatchProviderResults, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%stv?api_key=%s%s",
		baseURL,
		watchProvidersURL,
		c.apiKey,
		options,
	)
	watchProvider := MediaWatchProviderResults{}
	if err := c.get(tmdbURL, &watchProvider); err != nil {
		return nil, err
	}
	return &watchProvider, nil
}

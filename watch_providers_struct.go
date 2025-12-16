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
		ID                int64          `json:"id"`
		Name              string         `json:"name"`
		DisplayPriorities map[string]int `json:"display_priorities"`
		DisplayPriority   int64          `json:"display_priority"`
		LogoPath          string         `json:"logo_path"`
		ProviderName      string         `json:"provider_name"`
		ProviderID        int            `json:"provider_id"`
	} `json:"results"`
}

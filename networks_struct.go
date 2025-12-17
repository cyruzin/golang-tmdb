package tmdb

// NetworkDetails type is a struct for details JSON response.
type NetworkDetails struct {
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

// NetworkAlternativeNames type is a struct for alternative names JSON response.
type NetworkAlternativeNames struct {
	ID      int64 `json:"id"`
	Results []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"results"`
}

// NetworkImage type is a struct for a single image.
type NetworkImage struct {
	ImageBase
	ID       string `json:"id"`
	FileType string `json:"file_type"`
}

// NetworkImages type is a struct for images JSON response.
type NetworkImages struct {
	ID    int64          `json:"id"`
	Logos []NetworkImage `json:"logos"`
}

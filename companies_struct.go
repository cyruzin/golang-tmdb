package tmdb

// CompanyDetails type is a struct for details JSON response.
type CompanyDetails struct {
	Description   string `json:"description"`
	Headquarters  string `json:"headquarters"`
	Homepage      string `json:"homepage"`
	ID            int64  `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
	ParentCompany struct {
		Name     string `json:"name"`
		ID       int64  `json:"id"`
		LogoPath string `json:"logo_path"`
	} `json:"parent_company"`
}

// CompanyAlternativeNames type is a struct for alternative
// names JSON response.
type CompanyAlternativeNames struct {
	ID int64 `json:"id"`
	*CompanyAlternativeNamesResult
}

// CompanyImage type is a struct for a single image.
type CompanyImage struct {
	ImageBase
	ID       string `json:"id"`
	FileType string `json:"file_type"`
}

// CompanyImages type is a struct for images JSON response.
type CompanyImages struct {
	ID    int64          `json:"id"`
	Logos []CompanyImage `json:"logos"`
}

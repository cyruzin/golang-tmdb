package tmdb

// CompanyDetails type is a struct for details JSON response.
type CompanyDetails struct {
	CompanyInfoDetails
	Description   string       `json:"description"`
	ParentCompany *CompanyInfo `json:"parent_company"`
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

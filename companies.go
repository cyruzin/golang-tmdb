package tmdb

import (
	"fmt"
)

// CompanyDetails type is a struct for details JSON response.
type CompanyDetails struct {
	CompanyInfo
	Description   string       `json:"description"`
	Headquarters  string       `json:"headquarters"`
	Homepage      string       `json:"homepage"`
	ParentCompany *CompanyInfo `json:"parent_company"`
}

// GetCompanyDetails get a companies details by id.
//
// https://developers.themoviedb.org/3/companies/get-company-details
func (c *Client) GetCompanyDetails(
	id int,
) (*CompanyDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s",
		baseURL,
		companyURL,
		id,
		c.apiKey,
	)
	companyDetails := CompanyDetails{}
	if err := c.get(tmdbURL, &companyDetails); err != nil {
		return nil, err
	}
	return &companyDetails, nil
}

// GetCompanyAlternativeNames get the alternative names of a company.
//
// https://developer.themoviedb.org/reference/company-alternative-names
func (c *Client) GetCompanyAlternativeNames(
	id int,
) (*IDAlternativeNameResults, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_names?api_key=%s",
		baseURL,
		companyURL,
		id,
		c.apiKey,
	)
	companyAlternativeNames := IDAlternativeNameResults{}
	if err := c.get(tmdbURL, &companyAlternativeNames); err != nil {
		return nil, err
	}
	return &companyAlternativeNames, nil
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

// GetCompanyImages get a companies logos by id.
//
// There are two image formats that are supported for companies,
// PNG's and SVG's. You can see which type the original file is
// by looking at the file_type field. We prefer SVG's as they
// are resolution independent and as such, the width and height
// are only there to reflect the original asset that was uploaded.
// An SVG can be scaled properly beyond those dimensions if you
// call them as a PNG.
//
// https://developers.themoviedb.org/3/companies/get-company-images
func (c *Client) GetCompanyImages(
	id int,
) (*CompanyImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s",
		baseURL,
		companyURL,
		id,
		c.apiKey,
	)
	companyImages := CompanyImages{}
	if err := c.get(tmdbURL, &companyImages); err != nil {
		return nil, err
	}
	return &companyImages, nil
}

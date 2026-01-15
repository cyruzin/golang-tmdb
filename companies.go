package tmdb

import (
	"fmt"
)

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

// GetCompanyDetails get a companies details by id.
//
// https://developers.themoviedb.org/3/companies/get-company-details
func (c *Client) GetCompanyDetails(id int) (*CompanyDetails, error) {
	url := c.baseURL.JoinPath(companyPath, fmt.Sprintf("%d", id))
	var companyDetails CompanyDetails
	if err := c.get(url, nil, &companyDetails); err != nil {
		return nil, err
	}
	return &companyDetails, nil
}

// CompanyAlternativeNames type is a struct for alternative
// names JSON response.
type CompanyAlternativeNames struct {
	ID int64 `json:"id"`
	*CompanyAlternativeNamesResult
}

// GetCompanyAlternativeNames get the alternative names of a company.
//
// https://developers.themoviedb.org/3/companies/get-company-alternative-names
func (c *Client) GetCompanyAlternativeNames(id int) (*CompanyAlternativeNames, error) {
	url := c.baseURL.JoinPath(companyPath, fmt.Sprintf("%d", id), "alternative_names")
	var companyAlternativeNames CompanyAlternativeNames
	if err := c.get(url, nil, &companyAlternativeNames); err != nil {
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
func (c *Client) GetCompanyImages(id int) (*CompanyImages, error) {
	url := c.baseURL.JoinPath(companyPath, fmt.Sprintf("%d", id), "images")
	var companyImages CompanyImages
	if err := c.get(url, nil, &companyImages); err != nil {
		return nil, err
	}
	return &companyImages, nil
}

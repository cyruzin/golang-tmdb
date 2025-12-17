package tmdb

import (
	"fmt"
)

// GetCompanyDetails get a companies details by id.
//
// https://developer.themoviedb.org/reference/company-details
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
) (*CompanyAlternativeNames, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_names?api_key=%s",
		baseURL,
		companyURL,
		id,
		c.apiKey,
	)
	companyAlternativeNames := CompanyAlternativeNames{}
	if err := c.get(tmdbURL, &companyAlternativeNames); err != nil {
		return nil, err
	}
	return &companyAlternativeNames, nil
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
// https://developer.themoviedb.org/reference/company-images
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

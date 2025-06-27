package tmdb

import "fmt"

// GetNetworkDetails get the details of a network.
//
// https://developer.themoviedb.org/reference/network-details
func (c *Client) GetNetworkDetails(
	id int,
) (*CompanyInfoDetails, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s",
		baseURL,
		networkURL,
		id,
		c.apiKey,
	)
	networkDetails := CompanyInfoDetails{}
	if err := c.get(tmdbURL, &networkDetails); err != nil {
		return nil, err
	}
	return &networkDetails, nil
}

// GetNetworkAlternativeNames get the alternative names of a network.
//
// https://developer.themoviedb.org/reference/details-copy
func (c *Client) GetNetworkAlternativeNames(
	id int,
) (*IDAlternativeNameResults, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/alternative_names?api_key=%s",
		baseURL,
		networkURL,
		id,
		c.apiKey,
	)
	networkAltenativeNames := IDAlternativeNameResults{}
	if err := c.get(tmdbURL, &networkAltenativeNames); err != nil {
		return nil, err
	}
	return &networkAltenativeNames, nil
}

// GetNetworkImages get the TV network logos by id.
//
// There are two image formats that are supported for networks,
// PNG's and SVG's. You can see which type the original file is
// by looking at the file_type field. We prefer SVG's as they are
// resolution independent and as such, the width and height are
// only there to reflect the original asset that was uploaded.
// An SVG can be scaled properly beyond those dimensions if you
// call them as a PNG.
//
// https://developer.themoviedb.org/reference/alternative-names-copy
func (c *Client) GetNetworkImages(
	id int,
) (*CompanyImages, error) {
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s",
		baseURL,
		networkURL,
		id,
		c.apiKey,
	)
	networkImages := CompanyImages{}
	if err := c.get(tmdbURL, &networkImages); err != nil {
		return nil, err
	}
	return &networkImages, nil
}

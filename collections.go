package tmdb

import "fmt"

// GetCollectionDetails get collection details by id.
//
// https://developer.themoviedb.org/reference/collection-details
func (c *Client) GetCollectionDetails(
	id int,
	urlOptions map[string]string,
) (*CollectionDetails, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s",
		baseURL, collectionURL, id, c.apiKey, options,
	)
	collectionDetails := CollectionDetails{}
	if err := c.get(tmdbURL, &collectionDetails); err != nil {
		return nil, err
	}
	return &collectionDetails, nil
}

// GetCollectionImages get the images for a collection by id.
//
// https://developer.themoviedb.org/reference/collection-images
func (c *Client) GetCollectionImages(
	id int,
	urlOptions map[string]string,
) (*CollectionImages, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s%s",
		baseURL, collectionURL, id, c.apiKey, options,
	)
	collectionImages := CollectionImages{}
	if err := c.get(tmdbURL, &collectionImages); err != nil {
		return nil, err
	}
	return &collectionImages, nil
}

// GetCollectionTranslations get the list translations for a collection by id.
//
// https://developer.themoviedb.org/reference/collection-translations
func (c *Client) GetCollectionTranslations(
	id int,
	urlOptions map[string]string,
) (*CollectionTranslations, error) {
	options := c.fmtOptions(urlOptions)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s",
		baseURL, collectionURL, id, c.apiKey, options,
	)
	collectionTranslations := CollectionTranslations{}
	if err := c.get(tmdbURL, &collectionTranslations); err != nil {
		return nil, err
	}
	return &collectionTranslations, nil
}

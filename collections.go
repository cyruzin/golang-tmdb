package tmdb

import "fmt"

// CollectionDetails type is a struct for details JSON response.
type CollectionDetails struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
	Parts        []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIDs         []int64 `json:"genre_ids"`
		ID               int64   `json:"id"`
		MediaType        string  `json:"media_type"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		Popularity       float32 `json:"popularity"`
		VoteMetrics
	} `json:"parts"`
}

// GetCollectionDetails get collection details by id.
//
// https://developers.themoviedb.org/3/collections/get-collection-details
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

// CollectionImage type is a struct for a single image.
type CollectionImage struct {
	ImageBase
	Iso639_1 string `json:"iso_639_1"`
}

// CollectionImages type is a struct for images JSON response.
type CollectionImages struct {
	ID        int64             `json:"id"`
	Backdrops []CollectionImage `json:"backdrops"`
	Posters   []CollectionImage `json:"posters"`
}

// GetCollectionImages get the images for a collection by id.
//
// https://developers.themoviedb.org/3/collections/get-collection-images
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

// CollectionTranslations type is a struct for translations JSON response.
type CollectionTranslations struct {
	ID           int64         `json:"id"`
	Translations []Translation `json:"translations"`
}

// GetCollectionTranslations get the list translations for a collection by id.
//
// https://developers.themoviedb.org/3/collections/get-collection-translations
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

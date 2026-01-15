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
func (c *Client) GetCollectionDetails(id int, urlOptions map[string]string) (*CollectionDetails, error) {
	url := c.baseURL.JoinPath(collectionPath, fmt.Sprintf("%d", id))
	var collectionDetails CollectionDetails
	if err := c.get(url, urlOptions, &collectionDetails); err != nil {
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
func (c *Client) GetCollectionImages(id int, urlOptions map[string]string) (*CollectionImages, error) {
	url := c.baseURL.JoinPath(collectionPath, fmt.Sprintf("%d", id), "images")
	var collectionImages CollectionImages
	if err := c.get(url, urlOptions, &collectionImages); err != nil {
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
func (c *Client) GetCollectionTranslations(id int, urlOptions map[string]string) (*CollectionTranslations, error) {
	url := c.baseURL.JoinPath(collectionPath, fmt.Sprintf("%d", id), "translations")
	var collectionTranslations CollectionTranslations
	if err := c.get(url, urlOptions, &collectionTranslations); err != nil {
		return nil, err
	}
	return &collectionTranslations, nil
}

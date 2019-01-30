package tmdb

import "fmt"

// CollectionsDetails type is a struct for details JSON response.
type CollectionsDetails struct {
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
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float32 `json:"vote_average"`
		VoteCount        int64   `json:"vote_count"`
		Popularity       float32 `json:"popularity"`
	} `json:"parts"`
}

// CollectionsImages type is a struct for images JSON response.
type CollectionsImages struct {
	ID        int64 `json:"id"`
	Backdrops []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"backdrops"`
	Posters []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int64   `json:"vote_count"`
		Width       int     `json:"width"`
	} `json:"posters"`
}

// CollectionsTranslations type is a struct for translations JSON response.
type CollectionsTranslations struct {
	ID           int64 `json:"id"`
	Translations []struct {
		Iso3166_1   string `json:"iso_3166_1"`
		Iso639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title    string `json:"title"`
			Overview string `json:"overview"`
			Homepage string `json:"homepage"`
		} `json:"data"`
	} `json:"translations"`
}

// GetCollectionsDetails get collection details by id.
//
// https://developers.themoviedb.org/3/collections/get-collection-details
func (c *Client) GetCollectionsDetails(id int, o map[string]string) (*CollectionsDetails, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d?api_key=%s%s", baseURL, collectionURL, id, c.APIKey, options,
	)
	t := CollectionsDetails{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetCollectionsImages get the images for a collection by id.
//
// https://developers.themoviedb.org/3/collections/get-collection-images
func (c *Client) GetCollectionsImages(id int, o map[string]string) (*CollectionsImages, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/images?api_key=%s%s", baseURL, collectionURL, id, c.APIKey, options,
	)
	t := CollectionsImages{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// GetCollectionsTranslations get the list translations for a collection by id.
//
// https://developers.themoviedb.org/3/collections/get-collection-translations
func (c *Client) GetCollectionsTranslations(id int, o map[string]string) (*CollectionsTranslations, error) {
	options := c.fmtOptions(o)
	tmdbURL := fmt.Sprintf(
		"%s%s%d/translations?api_key=%s%s", baseURL, collectionURL, id, c.APIKey, options,
	)
	t := CollectionsTranslations{}
	err := c.get(tmdbURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

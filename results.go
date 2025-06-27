package tmdb

type ChangeMedia struct {
	ID    int64 `json:"id"`
	Adult bool  `json:"adult"`
}

type PaginatedChangesMediaResults struct {
	Results []ChangeMedia `json:"results"`
	PaginatedResultsMeta
}

type PaginatedCompanyInfoResults struct {
	Results []CompanyInfo `json:"results"`
	PaginatedResultsMeta
}

type PaginatedCollectionResults struct {
	Results []Collection `json:"results"`
	PaginatedResultsMeta
}

// type PaginatedPersonMediaResults struct {
// 	Results []PersonMedia `json:"results"`
// 	PaginatedResultsMeta
// }

type PaginatedPersonResults struct {
	Results []PersonResult `json:"results"`
	PaginatedResultsMeta
}

// MovieReleaseDatesResults Result Types
type MovieReleaseDateResults struct {
	Results []MovieReleaseDate `json:"results"`
}

type IDMovieReleaseDateResults struct {
	ID int64 `json:"id"`
	MovieReleaseDateResults
}

type WatchProvider struct {
	Link     string               `json:"link"`
	Free     *[]WatchProviderBase `json:"free"`
	Ads      *[]WatchProviderBase `json:"ads"`
	Flatrate *[]WatchProviderBase `json:"flatrate"`
	Rent     *[]WatchProviderBase `json:"rent"`
	Buy      *[]WatchProviderBase `json:"buy"`
}

type WatchProviderResults struct {
	Results map[string]WatchProvider `json:"results"`
}

type IDWatchProviderResults struct {
	ID int64 `json:"id"`
	WatchProviderResults
}

type PaginatedMovieListResults struct {
	Results []MovieList `json:"results"`
	PaginatedResultsMeta
}

type IDPaginatedMovieListResults struct {
	ID int64 `json:"id"`
	PaginatedMovieListResults
}

// TVContentRatingsResults Result Types
type TVContentRatingsResults struct {
	Results []ContentRatings `json:"results"`
}

type IDTVContentRatingsResults struct {
	ID int64 `json:"id"`
	TVContentRatingsResults
}

// TVEpisodeGroupsResults Result Types
type TVEpisodeGroupsResults struct {
	Results []TVEpisodeGroup `json:"results"`
}

type IDTVEpisodeGroupsResults struct {
	ID int64 `json:"id"`
	TVEpisodeGroupsResults
}

// KeywordsResults Result Types
type KeywordsResults struct {
	Results []Keyword `json:"results"`
}

type IDKeywordsResults struct {
	ID int64 `json:"id"`
	KeywordsResults
}

// TVScreenedTheatricallyResults Result Types
type TVScreenedTheatricallyResults struct {
	Results []ScreenedTheatrically `json:"results"`
}

type Video struct {
	ID          string `json:"id"`
	Iso639_1    string `json:"iso_639_1"`
	Iso3166_1   string `json:"iso_3166_1"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Official    bool   `json:"official"`
	PublishedAt string `json:"published_at"`
	Site        string `json:"site"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
}

type VideoResults struct {
	ID      int64   `json:"id"`
	Results []Video `json:"results"`
}

type PaginatedMovieMediaResults struct {
	Results []MovieMedia `json:"results"`
	PaginatedResultsMeta
}

type PaginatedMovieResults struct {
	Results []Movie `json:"results"`
	PaginatedResultsMeta
}

type PaginatedTVShowMediaResults struct {
	Results []TVShowMedia `json:"results"`
	PaginatedResultsMeta
}

type PaginatedTVShowResults struct {
	Results []TVShow `json:"results"`
	PaginatedResultsMeta
}

type PaginatedTVShowRatingResults struct {
	Results []TVShowRating `json:"results"`
	PaginatedResultsMeta
}

type PaginatedListResults struct {
	Results []List `json:"results"`
	PaginatedResultsMeta
}

type PaginatedMovieRatingResults struct {
	Results []MovieRating `json:"results"`
	PaginatedResultsMeta
}

type PaginatedTVEpisodeRatingResults struct {
	Results []TVEpisodeRating `json:"results"`
	PaginatedResultsMeta
}

type PaginatedMediaResults struct {
	Results []Media `json:"results"`
	PaginatedResultsMeta
}

type PaginatedKeywordsResults struct {
	KeywordsResults
	PaginatedResultsMeta
}

// KeywordMovies type is a struct for movies that belong to a keyword JSON response.
type IDPaginatedMovieResults struct {
	ID int64 `json:"id"`
	PaginatedMovieResults
}

type PaginatedReviewsResults struct {
	Results []Review `json:"results"`
	PaginatedResultsMeta
}

type IDPaginatedReviewsResults struct {
	ID int64 `json:"id"`
	PaginatedReviewsResults
}

type MediaWatchProviderResults struct {
	Results []MediaWatchProvider `json:"results"`
}

type CountryWatchProviderResults struct {
	Results []Country `json:"results"`
}

type AlternativeTitleResults struct {
	Results []AlternativeTitle `json:"results"`
}

type IDAlternativeTitleResults struct {
	ID int `json:"id"`
	AlternativeTitleResults
}

type AlternativeNameResults struct {
	Results []AlternativeName `json:"results"`
}

type IDAlternativeNameResults struct {
	ID int64 `json:"id"`
	AlternativeNameResults
}

type PaginatedDatesMovieResults struct {
	Dates Dates `json:"dates"`
	PaginatedMovieResults
}

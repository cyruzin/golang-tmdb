package tmdb

// SearchCompanies type is a struct for companies JSON response.
type SearchCompanies struct {
	*SearchCompaniesResults
	PaginatedResultsMeta
}

// SearchCollections type is a strcut for collections JSON response.
type SearchCollections struct {
	*SearchCollectionsResults
	PaginatedResultsMeta
}

// SearchKeywords type is a struct for keywords JSON response.
type SearchKeywords struct {
	*SearchKeywordsResults
	PaginatedResultsMeta
}

// SearchMovies type is a struct for movies JSON response.
type SearchMovies struct {
	PaginatedResultsMeta
	*SearchMoviesResults
}

// SearchMulti type is a struct for multi JSON response.
type SearchMulti struct {
	*SearchMultiResults
	PaginatedResultsMeta
}

// SearchPeople type is a struct for people JSON response.
type SearchPeople struct {
	PaginatedResultsMeta
	*SearchPeopleResults
}

// SearchTVShows type is a struct for tv show JSON response.
type SearchTVShows struct {
	PaginatedResultsMeta
	*SearchTVShowsResults
}

package tmdb

// DiscoverMovie type is a struct for movie JSON response.
type DiscoverMovie struct {
	PaginatedResultsMeta
	*DiscoverMovieResults
}

// DiscoverTV type is a struct for tv JSON response.
type DiscoverTV struct {
	PaginatedResultsMeta
	*DiscoverTVResults
}

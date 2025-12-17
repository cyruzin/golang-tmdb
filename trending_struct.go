package tmdb

// Trending type is a struct for trending JSON response.
type Trending struct {
	*TrendingResults
	PaginatedResultsMeta
}

package tmdb

type Dates struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}

type MovieListsDateResult struct {
	Dates Dates `json:"dates"`
	PaginatedMovieResults
}

package tmdb

// ChangesMovie type is a struct for movie changes JSON response.
type ChangesMovie struct {
	*ChangesMovieResults
	PaginatedResultsMeta
}

// ChangesTV type is a struct for tv changes JSON response.
type ChangesTV struct {
	*ChangesMovie
}

// ChangesPerson type is a struct for person changes JSON response.
type ChangesPerson struct {
	*ChangesMovie
}

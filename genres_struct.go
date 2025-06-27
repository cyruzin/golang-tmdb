package tmdb

// GenreMovieList type is a struct for genres movie list JSON response.
type GenreMovieList struct {
	Genres []Genre `json:"genres"`
}

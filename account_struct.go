package tmdb

type Gravatar struct {
	Hash string `json:"hash"`
}
type TMDB struct {
	AvatarPath *string `json:"avatar_path"`
}

type Avatar struct {
	Gravatar `json:"gravatar"`
	TMDB     `json:"tmdb"`
}

// AccountDetails type is a struct for details JSON response.
type AccountDetails struct {
	Avatar       Avatar `json:"avatar"`
	ID           int64  `json:"id"`
	Iso639_1     string `json:"iso_639_1"`
	Iso3166_1    string `json:"iso_3166_1"`
	Name         string `json:"name"`
	IncludeAdult bool   `json:"include_adult"`
	Username     string `json:"username"`
}

// AccountCreatedLists type is a struct for created lists JSON response.
type AccountCreatedLists struct {
	*AccountCreatedListsResults
	PaginatedResultsMeta
}

// AccountFavoriteMovies type is a struct for favorite movies JSON response.
type AccountFavoriteMovies struct {
	*AccountFavoriteMoviesResults
	PaginatedResultsMeta
}

// AccountFavoriteTVShows type is a struct for favorite tv shows JSON response.
type AccountFavoriteTVShows struct {
	*AccountFavoriteTVShowsResults
	PaginatedResultsMeta
}

// AccountFavorite type is a struct for movies or TV shows
// favorite JSON request.
type AccountFavorite struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Favorite  bool   `json:"favorite"`
}

// AccountRatedMovies type is a struct for rated movies JSON response.
type AccountRatedMovies struct {
	*AccountFavoriteMovies
}

// AccountRatedTVShows type is a struct for rated TV shows JSON response.
type AccountRatedTVShows struct {
	*AccountFavoriteTVShows
}

// AccountRatedTVEpisodes type is a struct for rated TV episodes JSON response.
type AccountRatedTVEpisodes struct {
	*AccountRatedTVEpisodesResults
	PaginatedResultsMeta
}

// AccountMovieWatchlist type is a struct for movie watchlist JSON response.
type AccountMovieWatchlist struct {
	*AccountFavoriteMovies
}

// AccountTVShowsWatchlist type is a struct for tv shows watchlist JSON response.
type AccountTVShowsWatchlist struct {
	*AccountFavoriteTVShows
}

// AccountWatchlist type is a struct for movies or TV shows
// watchlist JSON request.
type AccountWatchlist struct {
	MediaType string `json:"media_type"`
	MediaID   int64  `json:"media_id"`
	Watchlist bool   `json:"watchlist"`
}

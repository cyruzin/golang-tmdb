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

// AccountFavorite type is a struct for movies or TV shows
// favorite JSON request.
type AccountFavorite struct {
	MediaType MediaType `json:"media_type"`
	MediaID   int64     `json:"media_id"`
	Favorite  bool      `json:"favorite"`
}

// AccountWatchlist type is a struct for movies or TV shows
// watchlist JSON request.
type AccountWatchlist struct {
	MediaType MediaType `json:"media_type"`
	MediaID   int64     `json:"media_id"`
	Watchlist bool      `json:"watchlist"`
}
